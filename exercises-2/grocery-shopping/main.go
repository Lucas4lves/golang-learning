package main

import (
	"fmt"
	"shopping-list/models"
	"time"
)

func main() {
	watermelon := models.ShoppingItem{
		Name:  "Watermelon",
		Price: 1.99,
		Brand: "Nature",
	}

	list := models.ShoppingList{
		CreatedAt: time.Date(2024, 03, 28, 0, 0, 0, 0, time.Local),
		Shop:      "São Luiz",
		Items:     make([]models.ShoppingItem, 0),
	}

	list2 := models.ConstructShoppingList()

	list.AddShoppingItem(watermelon)
	list.AddShoppingItem(watermelon)
	list.AddShoppingItem(models.ShoppingItem{
		Name:  "Nintendo DS Lite",
		Price: 155.00,
		Brand: "Nintendo",
	})

	fmt.Println(list)
	fmt.Println(list2)
}
