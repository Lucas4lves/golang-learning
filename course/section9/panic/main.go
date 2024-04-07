package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./config_fie.txt")

	if err != nil {
		at_the_disco := fmt.Sprintf("Error: Config File reading failed - %s", err)
		panic(at_the_disco)
	}

	fmt.Println("Running state")
}

//panic halts the program
//its suited to prevent the execution if some critical error has happened
//e.g: database connection failure
