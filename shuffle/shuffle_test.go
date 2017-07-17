package shuffle

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	deck "github.com/robertkchang/deck_shuffle/deck"
)

func TestShuffle(t *testing.T) {
	// init new deck
	newDeck := deck.New()

	// make copy of new deck since shuffle will change newDeck
	var originalDeck = make([]deck.Card, 52)
	copy(originalDeck, newDeck)

	// shuffle
	shuffledDeck, _ := Shuffle(newDeck)

	// validate same number of cards returned
	assert.Equal(t, len(originalDeck), len(shuffledDeck), "Shuffled deck does not contain same number of cards as original deck")

	// validate first ten cards in shuffled deck not same as first ten cards in original
	numOfSameCards := 0
	for idx := 0; idx < 10; idx++ {
		if (originalDeck[idx].Suite == shuffledDeck[idx].Suite) &&
			(originalDeck[idx].Value == shuffledDeck[idx].Value) {
			numOfSameCards++
		}
	}
	assert.NotEqual(t, numOfSameCards, 10, "First ten cards in shuffled deck same as original deck")
}

func TestShuffleNilDeck(t *testing.T) {
	// shuffle with no deck
	_, error := Shuffle(nil)
	assert.EqualError(t, error, "deck.Card is missing or empty")
}

func TestShuffleEmptyDeck(t *testing.T) {
	// shuffle with no deck
	var emptyDeck []deck.Card
	_, error := Shuffle(emptyDeck)
	assert.EqualError(t, error, "deck.Card is missing or empty")
}

func TestSwap(t *testing.T) {
	// init new deck
	aDeck := deck.New()

	// make copy of new deck since shuffle will change newDeck
	var originalDeck = make([]deck.Card, 52)
	copy(originalDeck, aDeck)

	// save the original card 2 and 4 and swap
	cardTwo := originalDeck[2]
	cardFour := originalDeck[4]
	swap(aDeck, 2, 4)

	// validate that card 2 and 4 are swapped
	assert.Equal(t, cardTwo, aDeck[4], "Card 2 should be Card 4")
	assert.Equal(t, cardFour, aDeck[2], "Card 4 should be Card 2")
}

func TestGetRandomGen(t *testing.T) {
	// make ten random generators
	genArr := make([]*rand.Rand, 10)
	for idx := 0; idx < 10; idx++ {
		genArr[idx] = getRandomGen()
	}

	// validate none of the generators are the same
	numOfSameGen := 0
	for idx := 0; idx < 10; idx++ {
		// skip if 1st element
		if idx != 0 {
			if genArr[idx] == genArr[idx-1] {
				numOfSameGen++
			}
		}
	}
	assert.Equal(t, numOfSameGen, 0, "There were one or more random generators that are the same")
}
