package main

import (
	"fmt"
	"github.com/djordje200179/spellio"
	"github.com/djordje200179/spellio/layouts"
	"os"
)

func inputFromFile(fileName string, e *spellio.Engine) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	e.InsertFromText(file)
}

func main() {
	var engine spellio.Engine

	inputFromFile("books/test.txt", &engine)

	fmt.Printf("Words in dictionary: %d\n", engine.CountWords())

	fmt.Printf("hospi...: %v\n", engine.CompleteWord("hospi", 5))
	fmt.Printf("housr?: %v\n", engine.CorrectWord("housr", layouts.English, 5))
}
