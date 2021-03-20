package main

import (
	"fmt"
	"testing"
)

func TestNewCards (t * testing.T){ 
	d := newCards()

	if len(d) != 8{
		fmt.Println(len(d))
		fmt.Printf("Expected 8 but got %d as the length of the deck", len(d))
	}
}