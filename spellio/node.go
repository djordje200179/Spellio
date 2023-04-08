package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/linears/stack"
	"strings"
)

type WordInfo struct {
	Freq uint
}

type node struct {
	WordInfo
	char rune

	children []*node
	parent   *node
}

func (n *node) getChild(char rune) *node {
	for _, childNode := range n.children {
		if childNode.char == char {
			return childNode
		}
	}

	return nil
}

func (n *node) getWord(endNode *node) string {
	path := stack.New[*node]()
	for currNode := n; currNode != nil && currNode != endNode; currNode = currNode.parent {
		path.Push(currNode)
	}

	var sb strings.Builder
	path.ForEach(func(currNode *node) {
		sb.WriteRune(currNode.char)
	})

	return sb.String()
}
