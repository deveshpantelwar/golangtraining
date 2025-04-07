package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck lenght 16 but got ", len(d))
		//t.Errorf("Expected deck lenght 16 but got %v", len(d))

	}

	if d[0] != "Ace of Spades" {
		t.Error("expected 1st card ace of space but got ", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("expected Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Error("expected 16, but got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
