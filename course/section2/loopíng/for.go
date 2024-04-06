package main

import "fmt"

func main() {
	word := "Weird"
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}
}
