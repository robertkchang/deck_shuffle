package main

import (
  "fmt"

  deck "github.com/robertkchang/deck_shuffle/deck"
)

func main() {
  fmt.Println("=======================")

  // init
  original := deck.New()
  fmt.Println("[ORIGINAL] ", original)

  fmt.Println("=======================")

  // shuffle
  // NOTE: don't care about error in main() because Shuffle will always get a valid deck of cards
  shuffled, _ := deck.Shuffle(original)
  fmt.Println("[SHUFFLED] ", shuffled)
}
