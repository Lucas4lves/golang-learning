package main

import "fmt"

func multiOperations2(a int, b int) (int, int, int, int) {
	resultSum := a + b
	resultSub := a - b
	resultDiv := a / b
	resultMult := a * b

	return resultSum, resultSub, resultDiv, resultMult
}

func main() {
	sum, sub, div, mult := multiOperations2(4, 5)
	fmt.Println("All results: ", sum, sub, div, mult)
}
