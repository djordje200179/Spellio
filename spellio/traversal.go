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

		if currNode.WordInfo != nil {
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

func (e *Engine) GetWordsByPrefix(prefix string) map[string]WordInfo {
	prefix = strings.ToLower(prefix)

	startNode := e.findNode(prefix)

	nodeStack := stack.New[*letterNode]()
	nodeStack.Push(startNode)

	words := make(map[string]WordInfo)
	for !nodeStack.Empty() {
		currNode := nodeStack.Pop()

		if currNode.WordInfo != nil {
			word := prefix + currNode.getWord(startNode)
			words[word] = *currNode.WordInfo
		}

		for _, childNode := range currNode.children {
			nodeStack.Push(childNode)
		}
	}

	return words
}

type nearbyWordState struct {
	node    *letterNode
	changes uint
}

type NearbyWordInfo struct {
	WordInfo
	Changes uint
}

func (e *Engine) GetNearbyWords(rawWord string, maxChanges uint, layout KeyboardLayoutNearbyKeys) map[string]NearbyWordInfo {
	rawWord = strings.ToLower(rawWord)

	currStates := []nearbyWordState{{&e.root, 0}}
	for _, currChar := range rawWord {
		nextStates := make([]nearbyWordState, 0, len(currStates)*len(layout))

		for _, currState := range currStates {
			regularNextNode := currState.node.getChild(currChar)
			if regularNextNode != nil {
				regularNextState := nearbyWordState{regularNextNode, currState.changes}
				nextStates = append(nextStates, regularNextState)
			}

			if currState.changes < maxChanges {
				alternativeChars := layout[currChar]
				for _, alternativeChar := range alternativeChars {
					alternativeNextNode := currState.node.getChild(alternativeChar)
					if alternativeNextNode != nil {
						alternativeNextState := nearbyWordState{alternativeNextNode, currState.changes + 1}
						nextStates = append(nextStates, alternativeNextState)
					}
				}
			}
		}

		currStates = nextStates
	}

	possibleWords := make(map[string]NearbyWordInfo, len(currStates))
	for _, currState := range currStates {
		if currState.node.WordInfo == nil {
			continue
		}

		word := currState.node.getWord(&e.root)

		possibleWords[word] = NearbyWordInfo{*currState.node.WordInfo, currState.changes}
	}

	return possibleWords
}
