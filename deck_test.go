package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected the deck to have a length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected The first card to be Ace of Spades, instead %v is the first card", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, instead we got %v", d[len(d)-1])
	}
}

func TestOsFileNavigation(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected the deck to have a length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}