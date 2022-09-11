package main

import "fmt"

type Player interface {
	Cards() []Card
	Name() string
	Score() int
	GetCard(card Card)
}

type player struct {
	cards []Card
	name  string
	score int
}

func NewPlayer(name string) Player {
	return &player{name: name}
}

func (p *player) String() string {
	return fmt.Sprintf("Name: %s, Cards: %s, Score: %d \n", p.name, p.cards, p.score)
}

func (p *player) Cards() []Card {
	return p.cards
}

func (p *player) Name() string {
	return p.name
}

func (p *player) Score() int {
	withoutAce := []Card{}
	withAce := []Card{}
	totalWithoutAce := 0

	for _, c := range p.cards {
		if c.Rank == Ace {
			withAce = append(withAce, c)
		} else {
			totalWithoutAce += c.Value
			withoutAce = append(withoutAce, c)
		}
	}

	trackingTotal := totalWithoutAce

	for range withAce {
		totalWithAceAtEleven := trackingTotal + 11
		totalWithAceAtOne := trackingTotal + 1

		if totalWithAceAtEleven > 21 {
			trackingTotal = totalWithAceAtOne
		} else {
			trackingTotal = totalWithAceAtEleven
		}
	}

	p.score = trackingTotal

	return trackingTotal
}

func (p *player) GetCard(card Card) {
	p.cards = append(p.cards, card)
}
