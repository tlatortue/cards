package main

// create a new type of 'deck' which is a slice of strings

type deck []string

// function returns a value of type deck
func newDeck() deck {
	// new var of cards type deck
	// two slices. one of card suits, one of card values
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// replace unused variables (i & j) with _ in golang
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
}

// receiver function (var d type 'deck')
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}