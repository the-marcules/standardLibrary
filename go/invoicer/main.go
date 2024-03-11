package main

import (
	"bufio"
	"flag"
	"fmt"
	. "invoicer/company"
	. "invoicer/invoice"
	. "invoicer/person"
	"io"
	"os"
)

func setUpObjects() *Invoice {
	customer := NewRecipient(
		NewName("Herman", "Mayer", "Herr"),
		NewAddress("Altestrasse", "15", "50226", "Frechen", "Germany"),
	)

	dealer := NewCompany(
		"New Holland",
		"AG",
		"Brunsbuettel",
		NewAddress("BananaStr.", "1", "11111", "Brunsbuettel", "Germany"),
		NewPerson(
			NewName("David", "Hasselhoff", "Sir"),
			NewAddress("BongoBongo", "1", "01010", "Fort Laudadale", "USA"),
		),
		NewBilling("DE120000111144445555666777", "ASDFASDF11", "PenunsenBank"),
	)

	positions := GeneratePositions()

	invoice := NewInvoice(
		customer, dealer, positions, "€", 0.19,
	)

	return &invoice
}

func main() {
	var outputFlagPtr = flag.String("file", "", "if file is specified the invoice will be printed to file")
	var output io.Writer

	if *outputFlagPtr != "" {
		file, _ := os.Open(*outputFlagPtr)
		output = bufio.NewWriter(file)
		fmt.Printf("Invoice will be printed to file %s\n", *outputFlagPtr)
	} else {
		output = os.Stdout
		fmt.Println("Printing invoice on os.stdout")
	}

	invoice := setUpObjects()
	invoice.Print(output)
}
