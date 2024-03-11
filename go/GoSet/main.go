package main

import (
	"GoSet/pkg/set"
	"fmt"
)

type Person struct {
	name string
	age  int
	Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

func personStringer(p Person) string {
	if p.Country != "" {
		return fmt.Sprintf("%s (%d) aus %s %s", p.name, p.age, p.City, p.Country)
	}
	return fmt.Sprintf("%s (%d)", p.name, p.age)
}

func main() {
	p := Person{
		name: "De lange Tünn",
		age:  73,
		Address: Address{
			Street:  "Friesenstr.",
			City:    "Cologne",
			Country: "DE",
		},
	}
	travelGroup := set.NewSet(p, Person{name: "Heinz", age: 50, Address: Address{Country: "DE"}}, Person{name: "Hilde", age: 51, Address: Address{
		Street:  "Kohlengasse",
		City:    "Bottropp",
		Country: "DE",
	}}, Person{name: "Johnny", age: 28}, Person{name: "Purity", age: 21, Address: Address{Country: "USA", City: "Dallas, Texas"}}, p)
	travelGroup.Stringer = personStringer
	travelGroup.Add(p) // Tony should occur only once in the set
	println(travelGroup.String())
}
