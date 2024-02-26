package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	validLength := 30
	d := newDeck()

	if len(d) != validLength {
		t.Errorf("Expected deck length of %d, but got %v", validLength, (d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	validLength := 30
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != validLength {
		t.Errorf("Expected %d cards in deck, got %v", validLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
