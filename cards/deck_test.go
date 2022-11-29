package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	nd := newDeck()

	if len(nd) != 52 {
		t.Errorf("expected deck of size 52, got %d instead", len(nd))
	}

	if nd[0].toString() != "Ace of Spades" {
		t.Errorf("expected card %q, got %q instead", "Ace of Spades", nd[0])
	}

	if nd[51].toString() != "King of Clubs" {
		t.Errorf("expected card %q, got %q instead", "King of Clubs", nd[51])
	}
}

func TestSaveToFileLoadDeckFromFile(t *testing.T) {
	os.Remove("_testdeck")

	nd := newDeck()
	err := nd.saveToFile("_testdeck")
	if err != nil {
		t.Fatal(err)
	}

	ld := loadDeckFromFile("_testdeck")
	if len(ld) != 52 {
		t.Errorf("expected deck of size 52, got %d instead", len(nd))
	}
	os.Remove("_testdeck")
}
