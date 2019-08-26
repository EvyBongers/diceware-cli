package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
)

func main() {
	numberOfWords := flag.Int("number", 0, "Number of words to generate")
	delimiter := flag.String("delimiter", "", "Character to separate the words (may be left out for none)")
	flag.Parse()

	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(*numberOfWords)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Println(strings.Join(list, *delimiter))
}
