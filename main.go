package main

import (
	"fmt"
	"github.com/djordje200179/Spellio/spellio"
	"log"
	"os"
)

func main() {
	engine := spellio.Engine{}

	file, err := os.Open("sample_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	engine.AddManyWords(file)

	fmt.Println(engine.GetWordsByPrefix("ma"))

	engine.OutputAll(os.Stdout)
}
