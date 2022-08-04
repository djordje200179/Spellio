package spellio

type Node struct {
	freq uint
	char rune

	children []*Node
}

func (node *Node) getChild(char rune) *Node {
	for _, child := range node.children {
		if child.char == char {
			return child
		}
	}

	return nil
}
