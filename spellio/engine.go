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

func (e *Engine) AddWord(word string) {
	word = strings.ToLower(word)

	currNode := &e.root
	for _, char := range word {
		nextNode := currNode.getChild(char)

		if nextNode != nil {
			currNode = nextNode
		} else {
			newNode := &Node{char: char, parent: currNode}
			currNode.children = append(currNode.children, newNode)
			currNode = newNode
		}
	}

	currNode.freq++
}

func (e *Engine) AddManyWords(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		e.AddWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) IsWordValid(word string) bool {
	word = strings.ToLower(word)

	node := e.findNode(word)

	return node != nil
}

func (e *Engine) OutputAll(writer io.Writer) {
	for word, freq := range e.GetWordsByPrefix("") {
		fmt.Fprintf(writer, "%v (%v)\n", word, freq)
	}
}
