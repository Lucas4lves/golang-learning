package main

import (
	"fmt"
	"strconv"
)

// go routines - threads -> async execution
func main() {
	for x := 0; x < 10000; x++ {
		go PrintMessage(strconv.Itoa(x))
	}
}

func PrintMessage(msg string) {
	fmt.Println(msg)
}
