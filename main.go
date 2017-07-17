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
	// NOTE: don't care about error in main() because Shuffle will always get a valid deck of cards
	// TODO: would like the call to be deck.Shuffle
	shuffled, _ := shuffle.Shuffle(original)
	fmt.Println("[SHUFFLED] ", shuffled)
}
