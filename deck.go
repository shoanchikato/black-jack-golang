package main

type Deck interface {
	Cards() []Card
}

type deck struct{}

func NewDeck() Deck {
	return &deck{}
}

func (d *deck) Cards() []Card {
	return d.createFullDeck()
}

func (d *deck) createFullDeck() []Card {
	clubs := d.createQuarterDeck(Clubs)
	diamonds := d.createQuarterDeck(Diamonds)
	hearts := d.createQuarterDeck(Hearts)
	spades := d.createQuarterDeck(Spades)

	cards := []Card{}
	cards = append(cards, clubs...)
	cards = append(cards, diamonds...)
	cards = append(cards, hearts...)
	cards = append(cards, spades...)

	return cards
}

func (d *deck) createQuarterDeck(suit Suit) []Card {
	// create Jack, King, Queen cards
	JKQ := []Card{}
	for _, rank := range []Rank{Jack, King, Queen} {
		card := Card{suit, rank, 10}
		JKQ = append(JKQ, card)
	}

	// ace card
	Ace := Card{suit, Ace, 0}

	// cards 2 - 10
	TwoToTen := []Card{}
	for i := 1; i < 10; i++ {
		card := Card{suit, RankFromEnumValue(i), i + 1}
		TwoToTen = append(TwoToTen, card)
	}

	cards := []Card{}

	cards = append(cards, Ace)
	cards = append(cards, TwoToTen...)
	cards = append(cards, JKQ...)

	return cards
}
