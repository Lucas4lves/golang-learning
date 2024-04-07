package models

import "time"

type Person struct {
	Name      string
	Age       int
	NickName  string
	Address   Address
	BirthDate time.Time
}

func (p Person) CurrentAge() int {
	birthYear := p.BirthDate.Year()
	currentYear := time.Now().Year()
	return currentYear - birthYear
}

func GetAge(p Person) int {
	birthYear := p.BirthDate.Year()
	currentYear := time.Now().Year()
	return currentYear - birthYear
}

func (p *Person) ChangeAddress(newAddress Address) {
	p.Address = newAddress
}
