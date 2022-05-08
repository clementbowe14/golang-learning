package main

type Card struct {
	value int
	name  string
}

func createCard(value int, name string) *Card {
	card := Card{value: value, name: name}

	return &card
}

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {

// 	return "Five of Diamonds"
// }
