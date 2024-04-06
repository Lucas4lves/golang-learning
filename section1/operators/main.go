package main

import "fmt"

func main() {
	var grossIncome float32 = 1600
	var taxRate float32 = 0.10
	taxes := grossIncome * taxRate
	salary := grossIncome - taxes

	fmt.Printf("Salary: %f\n", salary)
	fmt.Printf("Taxes: %f\n", taxes)

}
