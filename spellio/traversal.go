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
