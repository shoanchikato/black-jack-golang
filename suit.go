package main

type Suit int

const (
	Clubs Suit = iota + 1
	Hearts
	Diamonds
	Spades
)

func (s Suit) String() string {
	return []string{"Clubs", "Hearts", "Diamonds", "Spades"}[s-1]
}
