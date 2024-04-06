package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value := "14.55"
	b, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal("Error while trying to parse bool")
	}
	fmt.Printf("Type: %T \n", b)
	fmt.Println(b)
}
