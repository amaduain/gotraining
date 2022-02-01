package main

import "fmt"

func main() {
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//hands, remainingDeck := deal(cards, 5)
	//hands.print()
	//fmt.Println("###########
	//cards.toFile("deck.txt")
	cards := newDeckFromFile("deck.txt")
	cards.print()
	fmt.Println("Shuffling")
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"

}
