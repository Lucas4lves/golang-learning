package main

import "fmt"

func main() {
	wholeList := []string{"Rice", "Beans", "Miso", "Curry", "Pork"}

	subList := wholeList[len(wholeList)-1:]

	fmt.Println(subList)
}
