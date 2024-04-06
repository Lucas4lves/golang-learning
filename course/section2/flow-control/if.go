package main

import "fmt"

func isUnderAge(age int) bool {
	if age >= 18 {
		return false
	}

	return true
}

func calculateSalary(grossSal float64) float64 {
	if grossSal >= 1000.00 {
		taxes := 0.10
		return grossSal - (grossSal * taxes)
	}

	return grossSal
}

func main() {
	res := calculateSalary(999.00)

	fmt.Print("RES: ")
	fmt.Printf("%f\n", res)
}
