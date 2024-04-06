package main

import "fmt"

func main() {
	res := map[string]int{
		"httpStatus": 200,
	}

	res["LOL"] = 4

	value, exists := res["LOL"]

	if exists {
		fmt.Println(value)
	}

	fmt.Println(res)

}
