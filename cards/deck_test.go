package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("The length of deck is expcected to be 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected the last card to be King of clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadeddeck := newDeckFromFile("_decktesting")

	if len(loadeddeck) != 52 {
		t.Errorf("Expecetd 52 cards in deck, but got %v", len(loadeddeck))
	}

	os.Remove("_decktesting")
}
