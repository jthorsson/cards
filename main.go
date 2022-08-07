package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("\n***************************")
	fmt.Println("Hand")
	fmt.Println("***************************")
	hand.print()
	fmt.Println("\n***************************")
	fmt.Println("Remaining Deck")
	fmt.Println("***************************")
	remainingDeck.print()

	cards.saveToFile("my_cards")

}
