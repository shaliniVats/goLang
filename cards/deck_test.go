package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Deck size must be 16, Actual deck size is %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 1st card of deck is Ace of Spades. got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of deck is Four of Clubs. got %v", len(d)-1)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Deck size must be 16, Actual deck size is %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
