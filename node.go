package spellio

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/collectionsequence"
	"strings"
)

type letterNode struct {
	*Word
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
	path := collectionsequence.NewDeque[*letterNode]()
	for currNode := node; currNode != nil && currNode != endNode; currNode = currNode.parent {
		path.PushBack(currNode)
	}

	var sb strings.Builder
	for !path.Empty() {
		currNode := path.PopBack()
		sb.WriteRune(currNode.char)
	}

	return sb.String()
}
