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

	curr := &e.root
	for _, char := range word {
		next, ok := curr.findChild(char)
		if !ok {
			next = &letterNode{char: char, parent: curr}
			curr.children = append(curr.children, next)
		}

		curr = next
	}

	if curr.Word == nil {
		curr.Word = &Word{
			Freq:       1,
			lastLetter: curr,
		}
	} else {
		curr.Freq++
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
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		word := sc.Text()
		word = cleanWord(word)

		e.Insert(word)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}

// FindWord will try to find a word in the Engine.
//
// If the word is found, the Word and true are returned.
// If the word is not found, an empty Word and false are returned.
func (e *Engine) FindWord(word string) (*Word, bool) {
	word = strings.ToLower(word)

	node, ok := e.findNode(word)
	if !ok {
		return nil, false
	}

	return node.Word, true
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
