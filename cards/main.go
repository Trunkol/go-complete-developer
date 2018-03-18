package main

func main() {
	cards := newDeck()
	hand, remainingOfDeck := deal(cards, 5)

	hand.print()
	remainingOfDeck.print()

}

func newCard() string {
	return "Five of Diamonds"
}
