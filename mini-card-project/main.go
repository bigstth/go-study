package main

func main() {
	cards := newDeckFromFile("my_cards.txt")

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	cards.shuffle()
	cards.saveToFile("my_cards.txt")
}
