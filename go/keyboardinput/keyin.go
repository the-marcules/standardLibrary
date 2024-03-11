package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	keyin := bufio.NewReader(os.Stdin)

DieUnendlicheMaschine:
	for {
		fmt.Println("What do you want to say? ")
		readString, err := keyin.ReadString('\n')
		if err != nil {
			fmt.Printf(err.Error())
			print("\n")
		} else {

			switch strings.TrimSuffix(readString, "\n") {
			case "/exit":
				break DieUnendlicheMaschine
			case "/q":
				break DieUnendlicheMaschine
			case "/quit":
				break DieUnendlicheMaschine
			default:
				fmt.Printf("Say: %s", readString)
			}

		}
	}
	print("\nBye!\n")

}
