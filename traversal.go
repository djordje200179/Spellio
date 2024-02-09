package spellio

import (
	"strings"
)

// CountWords returns the number of words in the dictionary.
func (e *Engine) CountWords() int {
	cnt := 0

	stack := []*letterNode{&e.root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Word != nil {
			cnt++
		}

		for _, child := range curr.children {
			stack = append(stack, child)
		}
	}

	return cnt
}

func (e *Engine) findNode(word string) (*letterNode, bool) {
	curr := &e.root
	for _, char := range word {
		next, ok := curr.findChild(char)
		if !ok {
			return nil, false
		}

		curr = next
	}

	return curr, true
}

// GetWordsByPrefix returns all words in the dictionary
// that start with the given prefix.
func (e *Engine) GetWordsByPrefix(prefix string) []*Word {
	prefix = strings.ToLower(prefix)
	start, ok := e.findNode(prefix)
	if !ok {
		return nil
	}

	var words []*Word

	stack := []*letterNode{start}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Word != nil {
			words = append(words, curr.Word)
		}

		for _, child := range curr.children {
			stack = append(stack, child)
		}
	}

	return words
}

// A NearbyWord contains the word and the number of changes
// required to transform it into the original word.
type NearbyWord struct {
	*Word
	Changes int
}

// A KeyboardLayout is a map of characters to
// characters that are nearby on the keyboard
// and are considered as possible replacements.
type KeyboardLayout map[rune][]rune

// GetNearbyWords returns all words in the dictionary
// that are near the given word and can be transformed
// into it with the given number of changes.
func (e *Engine) GetNearbyWords(rawWord string, maxChanges int, layout KeyboardLayout) []NearbyWord {
	rawWord = strings.ToLower(rawWord)
	rawWordChars := []rune(rawWord)

	possibleWords := make([]NearbyWord, 0)

	type nearbyWordState struct {
		node    *letterNode
		chars   []rune
		changes int
		index   int
	}

	queue := []nearbyWordState{{&e.root, []rune{}, 0, 0}}
	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.index == len(rawWordChars) {
			if currState.node.Word != nil {
				wordInfo := NearbyWord{currState.node.Word, currState.changes}
				possibleWords = append(possibleWords, wordInfo)
			}

			continue
		}

		nextChar := rawWordChars[currState.index]

		nextNode, ok := currState.node.findChild(nextChar)
		if ok {
			nextState := nearbyWordState{
				nextNode,
				make([]rune, len(currState.chars)+1),
				currState.changes + 1, currState.index + 1,
			}

			copy(nextState.chars, currState.chars)
			nextState.chars[len(currState.chars)] = nextChar

			queue = append(queue, nextState)
		}

		if currState.changes < maxChanges {
			alternativeChars := layout[nextChar]

			for _, alternativeChar := range alternativeChars {
				alternativeNextNode, ok := currState.node.findChild(alternativeChar)
				if !ok {
					continue
				}

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
