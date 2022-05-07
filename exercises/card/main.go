package main

import "fmt"

type Card struct {
	value int
	name  string
}

func createCard(value int, name string) *Card {
	card := Card{value: value, name: name}

	return &card
}
func main() {

	cards := deck{"king", "queen", "jack", "One of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for _, card := range cards {
		fmt.Println(card)
	}

	myFirstCard := createCard(14, "King of Hearts")
	fmt.Println(*myFirstCard)

}

func newCard() string {

	return "Five of Diamonds"
}
