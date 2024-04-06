package main

import "fmt"

func main() {
	x := 5
	y := &x

	fmt.Println(&x)
	fmt.Println(y)

	*y = 7

	fmt.Println(x)

}

func printValues(x int, y int) {
	fmt.Println(x)
	fmt.Println(y)
}
