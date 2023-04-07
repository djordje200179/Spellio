package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/linears/stack"
	"strings"
)

type WordInfo struct {
	Freq uint
}

type Node struct {
	WordInfo
	char rune

	children []*Node
	parent   *Node
}

func (node *Node) getChild(char rune) *Node {
	for _, childNode := range node.children {
		if childNode.char == char {
			return childNode
		}
	}

	return nil
}

func (node *Node) getWord(endNode *Node) string {
	path := stack.New[*Node]()
	for currNode := node; currNode != nil && currNode != endNode; currNode = currNode.parent {
		path.Push(currNode)
	}

	var sb strings.Builder
	path.ForEach(func(currNode *Node) {
		sb.WriteRune(currNode.char)
	})

	return sb.String()
}
