package main

func main() {
	cards := newDeck()

	// cards.shuffle().print()

	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.print()

	// hand.print()
	// leftover.print()

	// cards.saveToFile("cards.log")
	// hand.saveToFile("hand.log")
	// leftover.saveToFile("leftover.log")

	// test := loadDeckFromFile("hand.log")
	// test.print()
}
