package main

import (
	"fmt"

	deck "github.com/robertkchang/deck_shuffle/deck"
	shuffle "github.com/robertkchang/deck_shuffle/shuffle"
)

func main() {
	fmt.Println("=======================")

	// init
	original := deck.New()
	fmt.Println("[ORIGINAL] ", original)

	fmt.Println("=======================")

	// shuffle
	shuffled := shuffle.Shuffle(original)
	fmt.Println("[SHUFFLED] ", shuffled)
}
