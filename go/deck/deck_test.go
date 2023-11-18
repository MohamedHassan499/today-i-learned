package main

import "testing"

func TestNewCard(t *testing.T) {
	deck := newDeck()
	expectedLength := 16
	if len(deck) != expectedLength {
		t.Errorf("Expected deck of size %v, but got %v", expectedLength, len(deck))
	}
}

func TestFirstCard(t *testing.T) {
	deck := newDeck()
	expectedFirstCard := "Ace of Spades"
	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, deck[0])
	}
}

func TestLastCard(t *testing.T) {
	deck := newDeck()
	expectedLastCard := "Four of Clubs"
	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, deck[len(deck)-1])
	}
}