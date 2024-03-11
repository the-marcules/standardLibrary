package person

import (
	"fmt"
)

type Name struct {
	first string
	last  string
	title string
}

func NewName(first, last, title string) Name {
	return Name{
		first: first,
		last:  last,
		title: title,
	}

}

func (n Name) GetFullName() string {
	return fmt.Sprintf("%s %s %s", n.title, n.first, n.last)
}

type Address struct {
	Street      string
	houseNumber string
	postalCode  string
	city        string
	country     string
}

func NewAddress(street, houseNumber, postalCode, city, country string) Address {
	return Address{
		city:        city,
		houseNumber: houseNumber,
		postalCode:  postalCode,
		street:      street,
		country:     country,
	}
}

func (a Address) GetAddress() string {
	return fmt.Sprintf("%s %s\n%s %s\n%s", a.street, a.houseNumber, a.postalCode, a.city, a.country)
}

type Person struct {
	Name
	Address
}

func NewPerson(name Name, address Address) Person {
	return Person{
		name,
		address,
	}
}

func (p Person) GetFullAddress() string {
	return fmt.Sprintf("%s\n%s\n", p.GetFullName(), p.GetAddress())
}

type Recipient struct {
	Person
	CustomerID string
}

func NewRecipient(name Name, address Address) Recipient {

	return Recipient{
		Person: Person{
			Name:    name,
			Address: address,
		},
		CustomerID: "123",
	}

}
