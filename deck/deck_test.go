package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	newDeck := New()
	assert.Equal(t, len(newDeck), 52, "New deck should contain 52 cards")
}

func TestShuffleSameNumOfCards(t *testing.T) {
	// init new deck and shuffle
	originalDeck, shuffledDeck := createNewDeckAndShuffle()

	// validate same number of cards returned
	assert.Equal(t, len(originalDeck), len(shuffledDeck), "Shuffled deck does not contain same number of cards as original deck")
}

func TestShuffleValidateRandom(t *testing.T) {
	// init new deck and shuffle
	originalDeck, shuffledDeck := createNewDeckAndShuffle()

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

func TestShuffleValidateNoDuplicate(t *testing.T) {
	// init new deck and shuffle
	_, shuffledDeck := createNewDeckAndShuffle()

	// validate there are no duplicates after shuffle
	existMap := make(map[string]int)
	for _, card := range shuffledDeck {
		if _, ok := existMap[card.String()]; ok {
			// found a duplicate - fail the test immediately and break out of loop
			assert.Fail(t, "Duplicate found: ", card.String())
			break
		} else {
			// no duplicate - just increment counter in existMap
			existMap[card.String()]++
		}
	}
}

func TestShuffleNilDeck(t *testing.T) {
	// shuffle with no deck
	_, error := Shuffle(nil)
	assert.EqualError(t, error, "Card is missing or empty")
}

func TestShuffleEmptyDeck(t *testing.T) {
	// shuffle with empty deck
	var emptyDeck []Card
	_, error := Shuffle(emptyDeck)
	assert.EqualError(t, error, "Card is missing or empty")
}

func createNewDeckAndShuffle() ([]Card, []Card) {
	// init new deck
	newDeck := New()

	// make copy of new deck since shuffle will change newDeck
	var originalDeck = make([]Card, 52)
	copy(originalDeck, newDeck)

	// shuffle
	shuffledDeck, _ := Shuffle(newDeck)

	return originalDeck, shuffledDeck
}
