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

func (e *Engine) GetNearbyWords(rawWord string, maxDistance int, layout KeyboardLayoutNearbyKeys) map[string]uint {
	rawWord = strings.ToLower(rawWord)

	currNodes := []*Node{&e.root}
	for _, currChar := range rawWord {
		possibleChars := make([]rune, len(layout[currChar])+1)
		possibleChars[0] = currChar
		copy(possibleChars[1:], layout[currChar])

		nextNodes := make([]*Node, 0, len(possibleChars))
		for _, possibleChar := range possibleChars {
			for _, currNode := range currNodes {
				nextNode := currNode.getChild(possibleChar)
				if nextNode != nil {
					nextNodes = append(nextNodes, nextNode)
				}
			}
		}

		currNodes = nextNodes
	}

	possibleWords := make(map[string]uint, len(currNodes))
	for _, currNode := range currNodes {
		word := currNode.getWord(&e.root)
		occurrences := currNode.freq

		possibleWords[word] = occurrences
	}

	return possibleWords
}
