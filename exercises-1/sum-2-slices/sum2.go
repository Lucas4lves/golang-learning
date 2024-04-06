package main

import "fmt"

func getRangeFromASlice(s []int, low int, high int) []int {
	result := make([]int, 0)
	for k := 0; k < len(s); k++ {
		if s[k] >= low && s[k] <= high {
			result = append(result, s[k])
		}
	}

	return result
}

func sumWithinASlice(s []int) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result += s[i]
	}

	return result
}

func main() {
	fullSlice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	slice1 := getRangeFromASlice(fullSlice, 0, 5)
	slice2 := getRangeFromASlice(fullSlice, 6, 10)
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
	fmt.Println("Sum of slice1: ", sumWithinASlice(slice1))
	fmt.Println("Sum of slice2: ", sumWithinASlice(slice2))
}
