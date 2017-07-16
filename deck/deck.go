package deck

import (
  "fmt"
)

// Suites enum
var Suites = []string{
  "Spade",
  "Heart",
  "Club",
  "Diamond",
}

// Values of card enum
var Values = []string{
  "Ace",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
  "10",
  "Jack",
  "Queen",
  "King",
}

// Card struct
type Card struct {
  Value string
  Suite string
}

// Card's to String
func (card Card) String() string {
  return fmt.Sprintf("%s", "{"+card.Value+"|"+card.Suite+"}")
}

// New deck of cards
func New() []Card {
  deck := make([]Card, 52)

  var cardCnt = 0
  for _, suite := range Suites {
    for _, value := range Values {
      deck[cardCnt] = Card{Value: value, Suite: suite}
      cardCnt++
    }
  }

  return deck
}
