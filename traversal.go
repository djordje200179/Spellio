package spellio

import (
	"strings"
)

// CountWords returns the number of words in the dictionary.
func (e *Engine) CountWords() int {
	count := 0

	stack := []*letter{&e.root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Word != nil {
			count++
		}

		for _, child := range curr.children {
			stack = append(stack, child)
		}
	}

	return count
}

func (e *Engine) findNode(word string) *letter {
	curr := &e.root
	for _, char := range word {
		next := curr.getChild(char)

		if next != nil {
			curr = next
		} else {
			return nil
		}
	}

	return curr
}

// GetWordsByPrefix returns all words in the dictionary
// that start with the given prefix.
func (e *Engine) GetWordsByPrefix(prefix string) []Word {
	prefix = strings.ToLower(prefix)
	start := e.findNode(prefix)
	if start == nil {
		return nil
	}

	var words []Word

	stack := []*letter{start}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Word != nil {
			words = append(words, *curr.Word)
		}

		for _, child := range curr.children {
			stack = append(stack, child)
		}
	}

	return words
}

type nearbyWordState struct {
	node    *letter
	chars   []rune
	changes int
	index   int
}

// A NearbyWordInfo contains a word and the number of changes
// required to transform it into the original word.
type NearbyWordInfo struct {
	Word
	Changes int
}

// A KeyboardLayoutNearbyKeys is a map of characters
// to the characters that are nearby on the keyboard
// and are considered as possible replacements.
type KeyboardLayoutNearbyKeys map[rune][]rune

// GetNearbyWords returns all words in the dictionary
// that are near the given word and can be transformed
// into it with the given number of changes.
func (e *Engine) GetNearbyWords(rawWord string, maxChanges int, layout KeyboardLayoutNearbyKeys) []NearbyWordInfo {
	rawWord = strings.ToLower(rawWord)
	rawWordChars := []rune(rawWord)

	possibleWords := make([]NearbyWordInfo, 0)

	queue := []nearbyWordState{{&e.root, []rune{}, 0, 0}}
	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.index == len(rawWordChars) {
			if currState.node.Word != nil {
				wordInfo := NearbyWordInfo{*currState.node.Word, currState.changes}
				possibleWords = append(possibleWords, wordInfo)
			}

			continue
		}

		nextChar := rawWordChars[currState.index]

		regularNextNode := currState.node.getChild(nextChar)
		if regularNextNode != nil {
			nextStateChars := make([]rune, len(currState.chars)+1)
			copy(nextStateChars, currState.chars)
			nextStateChars[len(currState.chars)] = nextChar

			regularCharState := nearbyWordState{
				regularNextNode,
				//append(currState.chars, nextChar),
				nextStateChars,
				currState.changes, currState.index + 1,
			}

			queue = append(queue, regularCharState)
		}

		if currState.changes < maxChanges {
			alternativeChars := layout[nextChar]

			for _, alternativeChar := range alternativeChars {
				alternativeNextNode := currState.node.getChild(alternativeChar)

				if alternativeNextNode != nil {
					nextStateChars := make([]rune, len(currState.chars)+1)
					copy(nextStateChars, currState.chars)
					nextStateChars[len(currState.chars)] = alternativeChar

					alternativeCharState := nearbyWordState{
						alternativeNextNode,
						//append(currState.chars, alternativeChar),
						nextStateChars,
						currState.changes + 1, currState.index + 1,
					}

					queue = append(queue, alternativeCharState)
				}
			}

			var checkCharRedundancy bool = true
			// temporally every char will be checked for redundancy

			if checkCharRedundancy {
				nextStateChars := make([]rune, len(currState.chars))
				copy(nextStateChars, currState.chars)

				redundantCharState := nearbyWordState{
					currState.node,
					//currState.chars,
					nextStateChars,
					currState.changes + 1, currState.index + 1,
				}

				queue = append(queue, redundantCharState)
			}
		}
	}

	return possibleWords
}
