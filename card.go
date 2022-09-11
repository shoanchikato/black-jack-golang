package main

import "fmt"

type Card struct {
	Suit  Suit
	Rank  Rank
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("Suit: %s, Rank: %s, Value: %d \n", c.Suit, c.Rank, c.Value)
}

type CardQuery struct {
	Suit Suit
	Rank Rank
}
