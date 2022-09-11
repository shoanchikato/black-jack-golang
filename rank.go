package main

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	King
	Queen
	Jack
)

func (r Rank) String() string {
	return []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King ", "Queen", "Jack"}[r-1]
}

func RankFromEnumValue(n int) Rank {
	return []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, King, Queen, Jack}[n]
}
