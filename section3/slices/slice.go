package main

import "fmt"

func main() {
	a := []string{
		"Lucas",
		"Alves",
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
