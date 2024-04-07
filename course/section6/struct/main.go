package main

import "fmt"

func main() {
	fmt.Println("Structs!")

	u := Customer{
		1,
		"Lucas Alves",
		29,
		Address{"Rua 2", "-", "-"},
		make([]Order, 0),
	}

	u.orderHistory = append(u.orderHistory, Order{1, 1400.00, "A ball", "Ball"})

	fmt.Println(u)
}

type Address struct {
	StreetName string
	ZipCode    string
	Number     string
}

type Customer struct {
	id           int
	fullname     string
	age          int
	address      Address
	orderHistory []Order
}

type Order struct {
	orderId     int
	totalPrice  float64
	description string
	item        string
}
