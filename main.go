package main // main package should be involved in every file if it needs to run

import (
	"fmt"
	"os"
)

func main() {
	cards:= newCards()

	deal, remaining := cards.deal(4)

	err := writeToFile(deal, "deal.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err = writeToFile(remaining, "remaining.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	deal = toDeck(readFromFile("deal.txt")) 
	remaining = toDeck(readFromFile("remaining.txt")) 

	deal.print()
	remaining.print()
}