package main

import "fmt"

func main() {

	cards := []string{"king", "queen", "jack", "1", newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {

	return "Five of Diamonds"
}
