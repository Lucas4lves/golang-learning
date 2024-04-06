package main

import (
	"fmt"
	"log"
)

func findInAList(l []string, termToFind string) bool {
	if len(l) == 0 {

		log.Panic("The list is empty")
	}

	for x := 0; x < len(l); x++ {
		if l[x] == termToFind {
			fmt.Printf("Term found at %d of the list\n", x)
			return true
		}
	}

	return false
}

func main() {
	names := make([]string, 0)

	names = append(names, "Lucas")
	names = append(names, "Azvd")
	names = append(names, "Babner")
	names = append(names, "Jhon")
	names = append(names, "Mika")
	names = append(names, "Henso")
	names = append(names, "Xavi")

	if findInAList(names, "Henso") {
		fmt.Println("Success!")
	}

}
