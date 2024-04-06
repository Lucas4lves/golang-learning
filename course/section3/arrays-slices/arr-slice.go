package main

import "fmt"

/*Slices are ~like~ dynamic arrays*/
/*Arrays must have a predefined size*/
func main() {
	dummyArray := [3]int{10, 2, 78}
	dummmySlice := []int{5, 89, 7}

	fmt.Printf("Dummy: %T\n", dummyArray)
	fmt.Printf("Dummy: %T\n", dummmySlice)

}
