package main

import "fmt"

func printMessage(msg string) {
	fmt.Println(msg)
}

func setCharTitle(charName string, titleName string) string {
	newName := titleName + ": " + charName

	return newName
}

func main() {
	printMessage(setCharTitle("Lucas", "Sage"))
}
