package main

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Saved!")
		}
	}()

	_, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("File does not exist: %s", err))
	}
}

func main() {
	ReadFile("./not-exists.txt")

	fmt.Println("Finished")
}
