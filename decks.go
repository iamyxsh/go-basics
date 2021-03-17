package main

import (
	"fmt" //importing a package to print statements
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) deal(hand int) (deck,deck){ //returns 2 variables with type deck and deck
	return d[:hand], d[hand:]  //range of slice
}

func (d deck) toString() (string){
	return strings.Join(d, ",") //import from package string
}

func toDeck(s string) deck {
	return strings.Split(s, ",")
}

func writeToFile(d deck, fileName string) error { // returns a variable with type error
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) //import from package ioutil
} //write something to a file
// returns an error , error === nil when operation is a success

func readFromFile(fileName string) string {
	 d, err:= ioutil.ReadFile(fileName)
	 if err != nil{
		fmt.Println(err)
		 os.Exit(1)
	 }

	 return string(d)
}