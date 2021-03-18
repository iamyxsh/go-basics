package main

import (
	"fmt"
	"testing"
)

func TestNewCards (t * testing.T){ 
	d := newCards()

	if len(d) != 8{
		fmt.Println(len(d))
		t.Error("Expected 8 but got %d as the length of the deck", len(d))
	}
}