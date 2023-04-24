package spellio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

type Engine struct {
	root letterNode
}

func (e *Engine) Insert(word string) {
	word = strings.ToLower(word)

	currNode := &e.root
	for _, char := range word {
		nextNode := currNode.getChild(char)

		if nextNode != nil {
			currNode = nextNode
		} else {
			newNode := &letterNode{char: char, parent: currNode}
			currNode.children = append(currNode.children, newNode)
			currNode = newNode
		}
	}

	if currNode.Word == nil {
		currNode.Word = &Word{
			Freq:       1,
			lastLetter: currNode,
		}
	} else {
		currNode.Freq++
	}
}

var nonAlphaCharsRegex = regexp.MustCompile(`[^a-z]+`)

func cleanWord(word string) string {
	word = strings.ToLower(word)
	return nonAlphaCharsRegex.ReplaceAllString(word, "")
}

func (e *Engine) InsertFromText(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = cleanWord(word)

		e.Insert(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) FindWord(word string) (Word, bool) {
	word = strings.ToLower(word)

	node := e.findNode(word)

	if node == nil {
		return Word{}, false
	}

	return *node.Word, true
}

func (e *Engine) OutputAllWords(writer io.Writer) error {
	for _, word := range e.GetWordsByPrefix("") {
		_, err := fmt.Fprintf(writer, "%v (%v)\n", word, word.Freq)
		if err != nil {
			return err
		}
	}

	return nil
}
