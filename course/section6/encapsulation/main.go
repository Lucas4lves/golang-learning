package main

import (
	"fmt"
	"time"

	models "github.com/Lucas4lves/go-encapsulation/model"
)

func main() {
	a := models.Address{
		StreetName: "Rua 2",
		ZipCode:    "-",
		Number:     "-",
	}

	p := models.Person{
		Name:      "Lucas",
		Age:       29,
		NickName:  "Solid Snake",
		Address:   a,
		BirthDate: time.Date(1995, 03, 28, 0, 0, 0, 0, time.Local),
	}

	p.ChangeAddress(models.Address{
		StreetName: "RUA B",
		ZipCode:    "7128712128",
		Number:     "1829182",
	})

	fmt.Println(p)
}
