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

	inputFromFile("Lord of the rings 1.txt", &engine)
	inputFromFile("Lord of the rings 2.txt", &engine)
	inputFromFile("Lord of the rings 3.txt", &engine)

	fmt.Println(engine.GetNearbyWords("housr", 1, spellio.SerbianKeyboardLayout))
}
