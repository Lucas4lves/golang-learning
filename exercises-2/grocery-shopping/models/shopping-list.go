package models

import "time"

type ShoppingList struct {
	CreatedAt time.Time
	Shop      string
	Items     []ShoppingItem
}

func ConstructShoppingList() *ShoppingList {
	out := ShoppingList{}

	return &out
}

func (sl *ShoppingList) AddShoppingItem(i ShoppingItem) {
	sl.Items = append(sl.Items, i)
}
