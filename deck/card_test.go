package deck

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestString(t *testing.T) {
  aCard := Card{Value: "Ace", Suite: "Club"}
  assert.Equal(t, aCard.String(), "{Ace|Club}")
}
