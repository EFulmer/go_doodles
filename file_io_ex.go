package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	myFile, err := os.Create("go_file_test.txt")
	defer myFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	myFile.WriteString("test")

	logger := log.New(myFile, "log:", log.LstdFlags)

	logger.Println("logging")
}
