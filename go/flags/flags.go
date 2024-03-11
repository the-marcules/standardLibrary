package falgs

import (
	"flag"
	"fmt"
)

func main() {
	flagNamePtr := flag.String("name", "", "specifies the name")
	flagAgePtr := flag.Int("age", 1, "specifies the age")
	flag.Parse()

	fmt.Printf("You startet the app with name: %s \n", *flagNamePtr)
	fmt.Printf("You startet the app with age: %n", *flagAgePtr)
}
