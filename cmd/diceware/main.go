package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/sethvargo/go-diceware/diceware"
)

func main() {
	delimiters := "!#$%&()*+,-./:;<=>?@[\\]^_{|}0123456789"
	numberOfWords := flag.Int("number", 0, "Number of words to generate")
	flag.Parse()

	// Generate requested number of words using the diceware algorithm.
	list, err := diceware.Generate(*numberOfWords)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	for idx, word := range list {
		if idx > 0 {
			// Pick a delimiter from the list above
			fmt.Print(string(delimiters[rand.Intn(len(delimiters))]))
		}
		fmt.Print(word)
	}
	fmt.Println()
}
