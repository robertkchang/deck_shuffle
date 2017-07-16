package deck

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNew(t *testing.T) {
  newDeck := New()
  assert.Equal(t, len(newDeck), 52, "New deck should contain 52 cards")
}

func TestString(t *testing.T) {
  aCard := Card{Value: "Ace", Suite: "Club"}
  assert.Equal(t, aCard.String(), "{Ace|Club}")
}
