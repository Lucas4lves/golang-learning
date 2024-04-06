package main

import "fmt"

func main() {

	b := make([]int, 2)
	b = append(b, 3)
	b = append(b, 8)
	b = append(b, 7)
	b = append(b, 13)

	newList := make([]int, 0)

	for z := 0; z < len(b); z++ {
		fmt.Println(b[z])
		if b[z] < 10 {
			newList = append(newList, b[z])
		}
	}

	fmt.Printf("%T\nSize of: %d\n", b, len(b))
	fmt.Println(newList)
}
