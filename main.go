package main

import (
	"fmt"
	"github.com/djordje200179/Spellio/spellio"
	"log"
	"os"
)

func inputFromFile(fileName string, engine *spellio.Engine) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	engine.AddManyWords(file)
}

func main() {
	engine := spellio.Engine{}

	inputFromFile("books/Lord of the rings 1.txt", &engine)
	inputFromFile("books/Lord of the rings 2.txt", &engine)
	inputFromFile("books/Lord of the rings 3.txt", &engine)
	inputFromFile("books/Harry Potter - The Philosopher's Stone.txt", &engine)

	fmt.Println(engine.CountWords())
}
