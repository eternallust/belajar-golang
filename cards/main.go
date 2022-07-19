package main

func main() {
	// var card string = "Ace of Spades";
	// card := newCard();

	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()

	//
}