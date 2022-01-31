package main

import "fmt"

// Create a new type of 'deck'
// which is a slice os strings
type deck []string

// Create new deck, like an init method

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Print function for deck receiver

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
