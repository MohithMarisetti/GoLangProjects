package main

import "fmt"

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}
