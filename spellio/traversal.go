package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/linears/stack"
	"strings"
)

func (e *Engine) findNode(word string) *Node {
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

func (e *Engine) GetWordsByPrefix(prefix string) map[string]uint {
	prefix = strings.ToLower(prefix)

	startNode := e.findNode(prefix)

	nodeStack := stack.New[*Node]()
	nodeStack.Push(startNode)

	res := make(map[string]uint)
	for !nodeStack.Empty() {
		currNode := nodeStack.Pop()

		if currNode.freq > 0 {
			word := prefix + currNode.getWord(startNode)
			res[word] = currNode.freq
		}

		for _, childNode := range currNode.children {
			nodeStack.Push(childNode)
		}
	}

	return res
}

type nearbyWordState struct {
	node    *Node
	changes uint
}

func (e *Engine) GetNearbyWords(rawWord string, maxChanges uint, layout KeyboardLayoutNearbyKeys) map[string]uint {
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

	possibleWords := make(map[string]uint, len(currStates))
	for _, currState := range currStates {
		if currState.node.freq == 0 {
			continue
		}

		word := currState.node.getWord(&e.root)

		possibleWords[word] = currState.changes
	}

	return possibleWords
}
