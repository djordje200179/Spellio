package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/linears/stack"
	"strings"
)

func (e *Engine) CountWords() int {
	nodeStack := stack.New[*letterNode]()
	nodeStack.Push(&e.root)

	count := 0
	for !nodeStack.Empty() {
		currNode := nodeStack.Pop()

		if currNode.Word != nil {
			count++
		}

		for _, childNode := range currNode.children {
			nodeStack.Push(childNode)
		}
	}

	return count
}

func (e *Engine) findNode(word string) *letterNode {
	currNode := &e.root
	for _, char := range word {
		nextNode := currNode.getChild(char)

		if nextNode != nil {
			currNode = nextNode
		} else {
			return nil
		}
	}

	return currNode
}

func (e *Engine) GetWordsByPrefix(prefix string) map[string]Word {
	prefix = strings.ToLower(prefix)

	startNode := e.findNode(prefix)

	nodeStack := stack.New[*letterNode]()
	nodeStack.Push(startNode)

	words := make(map[string]Word)
	for !nodeStack.Empty() {
		currNode := nodeStack.Pop()

		if currNode.Word != nil {
			word := prefix + currNode.getWord(startNode)
			words[word] = *currNode.Word
		}

		for _, childNode := range currNode.children {
			nodeStack.Push(childNode)
		}
	}

	return words
}

type nearbyWordState struct {
	node    *letterNode
	chars   []rune
	changes int
	index   int
}

type NearbyWordInfo struct {
	Word
	Changes int
}

func (e *Engine) GetNearbyWords(rawWord string, maxChanges int, layout KeyboardLayoutNearbyKeys) map[string]NearbyWordInfo {
	rawWord = strings.ToLower(rawWord)
	rawWordChars := []rune(rawWord)

	possibleWords := make(map[string]NearbyWordInfo)

	statesQueue := []nearbyWordState{{&e.root, []rune{}, 0, 0}}

	for len(statesQueue) > 0 {
		currState := statesQueue[0]
		statesQueue = statesQueue[1:]

		if currState.index == len(rawWordChars) {
			if currState.node.Word != nil {
				word := string(currState.chars)
				possibleWords[word] = NearbyWordInfo{*currState.node.Word, currState.changes}
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

			statesQueue = append(statesQueue, regularCharState)
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

					statesQueue = append(statesQueue, alternativeCharState)
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

				statesQueue = append(statesQueue, redundantCharState)
			}
		}
	}

	return possibleWords
}
