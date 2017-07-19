package deck

import (
	"errors"
	"math/rand"
	"time"
)

// New deck of cards
func New() []Card {
	deck := make([]Card, 52)

	var cardCnt = 0
	for _, suite := range Suites {
		for _, value := range Values {
			deck[cardCnt] = Card{Suite: suite, Value: value}
			cardCnt++
		}
	}

	return deck
}

// Shuffle the deck
// NOTE: slice argument is a header which contains pointers to the backing array
// thus, when the slice is updated, the array values are updated
// https://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value
func Shuffle(arr []Card) ([]Card, error) {
	// check if deck is nil or empty
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("Card is missing or empty")
	}

	// get random generator with new seed
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// using Durstenfeld variation of the Fisher-Yates Shuffle algorithm
	// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	for idx := len(arr) - 1; idx > 0; idx-- {
		// random int 0 <= rndNum < idx
		rndNum := randomGen.Intn(idx)

		// swap idx and rndNum
		arr[idx], arr[rndNum] = arr[rndNum], arr[idx]
	}
	return arr, nil
}
