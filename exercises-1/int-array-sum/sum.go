package main

import "fmt"

/*
Create an array of integers with 2 indexes,
sum them all in a variable and print it to console
*/
func main() {
	arr := [2]int{5, 4}
	sum := 0
	for y := 0; y < len(arr); y++ {
		sum += arr[y]
	}

	fmt.Println("Sum: ", sum)
}
