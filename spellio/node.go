package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/linears/stack"
	"strings"
)

type WordInfo struct {
	Freq uint
}

type letterNode struct {
	WordInfo
	char rune

	children []*letterNode
	parent   *letterNode
}

func (node *letterNode) getChild(char rune) *letterNode {
	for _, childNode := range node.children {
		if childNode.char == char {
			return childNode
		}
	}

	return nil
}

func (node *letterNode) getWord(endNode *letterNode) string {
	path := stack.New[*letterNode]()
	for currNode := node; currNode != nil && currNode != endNode; currNode = currNode.parent {
		path.Push(currNode)
	}

	var sb strings.Builder
	path.ForEach(func(currNode *letterNode) {
		sb.WriteRune(currNode.char)
	})

	return sb.String()
}
