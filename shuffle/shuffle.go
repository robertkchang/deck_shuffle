package shuffle

import (
	"errors"
	"math/rand"
	"time"

	deck "github.com/robertkchang/deck_shuffle/deck"
)

// Shuffle the deck
// NOTE: slice argument is a header which contains pointers to the backing array
// thus, when the slice is updated, the array values are updated
// https://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value
func Shuffle(arr []deck.Card) ([]deck.Card, error) {
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("deck.Card is missing or empty")
	}

	// get random generator with new seed
	randomGen := getRandomGen()

	// using Durstenfeld variation of the Fisher-Yates Shuffle algorithm
	// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	for idx := len(arr) - 1; idx > 0; idx-- {
		// random int 0 <= rndNum < idx
		rndNum := randomGen.Intn(idx)

		// swap idx and rndNum
		arr = swap(arr, idx, rndNum)
	}
	return arr, nil
}

// get random generator
func getRandomGen() *rand.Rand {
	// set random seed to current time
	seed := time.Now().UnixNano()
	seedSource := rand.NewSource(seed)
	return rand.New(seedSource)
}

// swap two array elements
func swap(arr []deck.Card, firstIdx int, secondIdx int) []deck.Card {
	tmp := arr[firstIdx]
	arr[firstIdx] = arr[secondIdx]
	arr[secondIdx] = tmp
	return arr
}
