package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// fmt.Println(card)

	// cards := []string{"Ace of Spades", newCard()}
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()
	cards.saveToFile("my_cards")

	fmt.Println("*****************")

	// cards.print()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("*****************")
	remainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func newCard() int {
// 	return 12
// }
