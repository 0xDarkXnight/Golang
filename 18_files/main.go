package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This needs to go in a file - https://github.com/0xDarkXnight"

	file, err := os.Create("./darkxnight.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("The length is:", length)

	defer file.Close()
	readFile("./darkxnight.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is:", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
