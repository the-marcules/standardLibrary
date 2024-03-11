package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MME struct {
	AirHumidity string `attribute:"0/--23-RIA" meta:"dasist1"`
	DingDong    string `attribute:"1-2-23-GNID" meta:"dasist1"`
	AffenGeil   string `attribute:"0-0-0-NEFFA" meta:"dasist2"`
}

func main() {
	mme := MME{
		AirHumidity: "20",
	}
	fmt.Println("Air ", mme.AirHumidity)

	rf := reflect.TypeOf(mme)

	for i := 0; i < rf.NumField(); i++ {
		feld := rf.Field(i)
		fmt.Println("Name: " + feld.Name)
		fmt.Println("Tag: ", feld.Tag)

		tags := strings.Split(fmt.Sprintf("%s", feld.Tag), " ")
		for _, tagPairStr := range tags {
			keyValArr := strings.SplitN(tagPairStr, ":", 2)
			fmt.Printf("- %s = %s", keyValArr[0], keyValArr[1][1:len(keyValArr[1])-1])
			println("")
		}

		/*		fmt.Println("Tag: ", feld.Tag.Get("attribute"))
				fmt.Println("Tag: ", feld.Tag.Get("meta"))*/
		fmt.Println("")
	}

}
