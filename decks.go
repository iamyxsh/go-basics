package main

import "fmt" //importing a package to print statements

type deck []string // declaring decks that extends the property of slice of type strings

func newCards() deck {
	cards := deck{} //making cards as the same data type of deck
	cardsSuit := []string{"Diamonds", "Hearts"} 
	cardsNumber := []string{"Ace", "Two", "Three", "Four"}

	//running loops to create a new deck
	for _, suit := range cardsSuit {  //same as js where suit is a single elemnt of cardsSuit
		for _, num := range cardsNumber { // range is equal to map
			cards = append(cards, num+ " of "+suit)
		}
	}

	return cards
}

func (d deck) print(){ // d is a reciever, same as this in js
	for i, card := range d{ // every elemnet with type deck can use this func
		fmt.Println(i, card)
	}
}