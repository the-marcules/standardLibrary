package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type Address struct {
	gorm.Model
	Street  string
	Number  string
	Zip     string
	City    string
	Country string
}
type Person struct {
	gorm.Model
	First     string
	Last      string
	Title     string
	AddressID uint
	Address   Address
}

func main() {
	db, err := gorm.Open(mysql.Open("root:rootpw@tcp(127.0.0.1:3306)/gosql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{}, &Person{}, &Address{})
	if err != nil {
		return
	}

	//result := db.First(&Person{}, 2)
	db.Unscoped().Delete(&Person{}, 1)
	var allPersons []Person
	db.Preload(clause.Associations).Find(&allPersons)

	for _, person := range allPersons {

		jsonStr, _ := json.Marshal(person)

		fmt.Printf("%s\n", string(jsonStr))
	}

	/*	markus := Person{First: "Markus", Last: "Fischer", Address: Address{Street: "Amselweg", Number: "16", Zip: "85111", City: "Adelschlag", Country: "Germany"}}
		//// Create
		db.Create(&markus)
	*/
}
