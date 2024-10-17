package main

import "fmt"

type deck [] string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamounds", "hearts", "clubs"}
	cardNumbers := []string{"ace", "two", "three","four", "five"}

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}