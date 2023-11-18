package main

func main() {
	deck := newDeck()	
	deck.shuffle()
	deck.print()
}

func newCard() string {
	return "Five of Diamonds"
}