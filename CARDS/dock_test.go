package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length of 16, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("First card is not correct (Ace of Spades) but got", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Error("Last card is not correct (Four of Clubs) but got", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckFile := "_decktesting"
	os.Remove(deckFile)
	d := newDeck()
	d.toFile(deckFile)
	loadedDeck := newDeckFromFile(deckFile)
	if len(loadedDeck) != 16 {
		t.Error("Expected deck length of 16, but got ", len(loadedDeck))
	}
	os.Remove(deckFile)
}
