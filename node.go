package spellio

import (
	"strings"
)

type letter struct {
	*Word
	char rune

	children []*letter
	parent   *letter
}

func (l *letter) getChild(char rune) *letter {
	for _, child := range l.children {
		if child.char == char {
			return child
		}
	}

	return nil
}

func (l *letter) getWord(end *letter) string {
	var path []*letter
	for curr := l; curr != nil && curr != end; curr = curr.parent {
		path = append(path, curr)
	}

	var sb strings.Builder
	for i := len(path) - 1; i >= 0; i-- {
		sb.WriteRune(path[i].char)
	}

	return sb.String()
}
