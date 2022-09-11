package main

import (
	"fmt"
)

func main() {
	deck := NewDeck()

	players := []Player{
		NewPlayer("Host"),
		NewPlayer("Tanya"),
		NewPlayer("Adam"),
		NewPlayer("Cain"),
		NewPlayer("Michelle"),
	}

	blackJack := NewBlackJack(deck, players)

	blackJack.HitPlayer("Tanya", CardQuery{Spades, Ace})
	blackJack.HitPlayer("Tanya", CardQuery{Hearts, Seven})
	blackJack.HitPlayer("Tanya", CardQuery{Diamonds, Ace})

	blackJack.HitPlayer("Adam", CardQuery{Diamonds, King})
	blackJack.HitPlayer("Adam", CardQuery{Spades, Four})
	blackJack.HitPlayer("Adam", CardQuery{Clubs, Four})

	blackJack.HitPlayer("Cain", CardQuery{Spades, Two})
	blackJack.HitPlayer("Cain", CardQuery{Diamonds, Two})
	blackJack.HitPlayer("Cain", CardQuery{Hearts, Two})
	blackJack.HitPlayer("Cain", CardQuery{Diamonds, Four})
	blackJack.HitPlayer("Cain", CardQuery{Clubs, Five})

	blackJack.HitPlayer("Michelle", CardQuery{Clubs, Queen})
	blackJack.HitPlayer("Michelle", CardQuery{Spades, Six})
	blackJack.HitPlayer("Michelle", CardQuery{Diamonds, Nine})

	blackJack.HitPlayer("Host", CardQuery{Spades, Jack})
	blackJack.HitPlayer("Host", CardQuery{Hearts, Nine})

	blackJack.RoundComplete()
	blackJack.GetScores()

	fmt.Println("Number of cards: ", len(blackJack.Cards()))

	for _, p := range blackJack.Winners() {
		fmt.Println("Winner: \n", p)
	}

	for _, p := range blackJack.Losers() {
		fmt.Println("Loser: \n", p)
	}

	for _, p := range blackJack.Ties() {
		fmt.Println("Tie: \n", p)
	}
}
