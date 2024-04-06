package main

import "fmt"

func multiOperations2(a int, b int) (sum int, sub int, div int, mult int) {
	sum = a + b
	sub = a - b
	div = a / b
	mult = a * b

	return
}

func main() {
	sm, sb, dv, mt := multiOperations2(8, 9)

	fmt.Println(sm, sb, dv, mt)
}
