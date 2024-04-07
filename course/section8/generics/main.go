package main

import "fmt"

func main() {
	fmt.Println("Generics")
	slice := []int{1, 2, 3}
	slice2 := []string{
		"Arroz",
		"Feijao",
		"Macarrao",
	}
	revertedSlice := RevertSlice(slice)

	fmt.Println(revertedSlice)

	revertedSlice2 := RevertSlice(slice2)

	fmt.Println(revertedSlice2)

	/*
		Output
		[3 2 1]
		[Macarrao Feijao Arroz]
	*/
}

/*Walking/Crawling pointers are tied to data structures specs,
not data types*/

//Generics are able to solve code duplication when its related solely to Data types

type CustomConstraint interface {
	int | string
}

func RevertSlice[T CustomConstraint](sl []T) []T {
	sliceCopy := make([]T, len(sl))
	walkingPointer := 0

	for k := len(sl) - 1; k >= 0; k-- {
		if walkingPointer < len(sl) {
			sliceCopy[walkingPointer] = sl[k]
		}
		walkingPointer++
	}

	return sliceCopy
}
