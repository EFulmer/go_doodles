package main

import (
	"bytes"
	"log"
	// "os"
)

const FILE_NAME = "example.log"

func main() {
	var out bytes.Buffer
	logger := log.New(&out, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")
	logger.Fatal("nooo")
}
