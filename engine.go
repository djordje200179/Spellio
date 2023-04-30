// Package spellio provides a trie structure for storing
// words and their frequencies and later spell checking
// and correction.
//
// The trie is optimized for fast lookup and traversal.
// It is also optimized for space efficiency.
package spellio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

// An Engine is a trie structure that stores information
// about words and their frequencies.
//
// Zero value is a valid empty Engine.
type Engine struct {
	root letterNode
}

// Insert adds a word to the Engine.
// If the word already exists, its frequency is incremented.
//
// The word is converted to lowercase before being inserted.
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

// InsertFromText reads words from a reader and inserts
// them into the Engine.
//
// Words will be separated by a space character, cleaned
// of non-alphabetic characters and converted to lowercase
// before being inserted.
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

// FindWord will try to find a word in the Engine.
//
// If the word is found, the Word and true are returned.
// If the word is not found, an empty Word and false are returned.
func (e *Engine) FindWord(word string) (Word, bool) {
	word = strings.ToLower(word)

	node := e.findNode(word)

	if node == nil {
		return Word{}, false
	}

	return *node.Word, true
}

// OutputAllWords will write all words in the Engine to a writer.
//
// Each word will be on its own line, with its frequency in parentheses.
func (e *Engine) OutputAllWords(writer io.Writer) error {
	for _, word := range e.GetWordsByPrefix("") {
		_, err := fmt.Fprintf(writer, "%v (%v)\n", word, word.Freq)
		if err != nil {
			return err
		}
	}

	return nil
}
