package spellio

import (
	"slices"
	"strings"
)

type letterNode struct {
	*Word
	char rune

	children []*letterNode
	parent   *letterNode
}

func (l *letterNode) findChild(char rune) (*letterNode, bool) {
	for _, child := range l.children {
		if child.char == char {
			return child, true
		}
	}

	return nil, false
}

func (l *letterNode) getWord(end *letterNode) string {
	var path []*letterNode
	for curr := l; curr != nil && curr != end; curr = curr.parent {
		path = append(path, curr)
	}

	slices.Reverse(path)

	var sb strings.Builder
	for _, letter := range path {
		sb.WriteRune(letter.char)
	}

	return sb.String()
}
