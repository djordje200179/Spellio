package spellio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Engine struct {
	root Node
}

func (e *Engine) Add(word string) {
	word = strings.ToLower(word)

	node := &e.root
	for _, char := range word {
		next := node.getChild(char)

		if next != nil {
			node = next
		} else {
			newNode := &Node{char: char}
			node.children = append(node.children, newNode)
			node = newNode
		}
	}

	node.freq++
}

func (e *Engine) IsValid(word string) bool {
	word = strings.ToLower(word)

	node := e.getNode(word)

	return node != nil
}

func recursiveTrieTraversal(node *Node, res map[string]uint, current string) {
	if node.freq > 0 {
		res[current] = node.freq
	}

	for _, child := range node.children {
		childString := current + string(child.char)
		recursiveTrieTraversal(child, res, childString)
	}
}

func (e *Engine) GetWordsWithPrefix(prefix string) map[string]uint {
	prefix = strings.ToLower(prefix)
	node := e.getNode(prefix)

	freqMap := make(map[string]uint)
	recursiveTrieTraversal(node, freqMap, prefix)

	return freqMap
}

func (e *Engine) Input(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		e.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Output(writer io.Writer) {
	for word, freq := range e.GetWordsWithPrefix("") {
		fmt.Fprintf(writer, "%v (%v)\n", word, freq)
	}
}

func (e *Engine) getNode(word string) *Node {
	node := &e.root
	for _, char := range word {
		next := node.getChild(char)

		if next != nil {
			node = next
		} else {
			return nil
		}
	}

	return node
}
