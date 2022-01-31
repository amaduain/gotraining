package main

import "fmt"

func main() {
	cards := newDeck()
	hands, remainingDeck := deal(cards, 5)
	hands.print()
	fmt.Println("################ REMAINING DECK ###############")
	remainingDeck.print()

}

func newCard() string {
	return "Five of Diamonds"

}
