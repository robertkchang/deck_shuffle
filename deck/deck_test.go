package deck

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNew(t *testing.T) {
  newDeck := New()
  assert.Equal(t, len(newDeck), 52, "New deck should contain 52 cards")
}

func TestShuffle(t *testing.T) {
  // init new deck
  newDeck := New()

  // make copy of new deck since shuffle will change newDeck
  var originalDeck = make([]Card, 52)
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
  // shuffle with empty deck
  var emptyDeck []Card
  _, error := Shuffle(emptyDeck)
  assert.EqualError(t, error, "deck.Card is missing or empty")
}
