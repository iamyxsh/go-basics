package main // main package should be involved in every file if it needs to run

import (
	"fmt"
)

func main() {
	cards:= newCards()

	cards.shuffle()

	fmt.Println(cards)

	
}