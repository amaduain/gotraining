package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Deal deck function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Print function for deck receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Convert the Deck to a list of strings

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Print the deck to file.
func (d deck) toFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

//Shuffle function
func (d deck) shuffle() {
	//calculate seed :-D
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	l := len(d) - 1
	for i := range d {
		newPosition := r.Intn(l)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
