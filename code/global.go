package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 5)

func main() {
	logger.Println("Hello, playground")
	// or
	log.Println("Hello, playground")
}
