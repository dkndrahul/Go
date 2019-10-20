package main

func main() {
	// var card string = "Ace of spades"
	// or
	// card := "Ace of spades" // 100 % equivalent to the above declaration
	// card = "Five of Diamonds"
	// fmt.Println(card)
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
