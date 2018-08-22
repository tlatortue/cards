package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// only use := for new variables
	// card := newCard()

	// array = fixed length, slice can grow / strink (must be same type)
	//cards := []string{"blah", "blabla"}

	cards := deck{"Ace of Diamonds", newCard()}
	// append
	cards = append(cards, "Six of Spades")

	// For loops
	for i, card:= range cards {
		fmt.Println(i, card)
	}
}


func newCard() string {
	return "Five of Diamonds"
}