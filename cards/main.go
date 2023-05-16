package main

func main() {
	// cards := newDeckFromFile("my_cards")
	// // hand, remainingCards := deal(cards, 5)
	// // hand.printDeck()
	// // remainingCards.printDeck()
	// cards.shuffle()
	// cards.printDeck()

	nums := []int{5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range nums {
		if i%2 == 0 {
			println(i, "is even.")
		} else {
			println(i, "is odd.")
		}
	}
}
