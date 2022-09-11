package main

import "fmt"

type BlackJack interface {
	Cards() []Card
	Players() []Player
	Winners() []Player
	Losers() []Player
	Ties() []Player
	HitPlayer(playerName string, card CardQuery)
	GetScores()
	RoundComplete()
}

type blackJack struct {
	cards   []Card
	players []Player
	winners []Player
	ties    []Player
	losers  []Player
}

func NewBlackJack(deck Deck, players []Player) BlackJack {
	cards := deck.Cards()

	return &blackJack{
		players: players,
		cards:   cards,
	}
}

func (b *blackJack) Cards() []Card {
	return b.cards
}

func (b *blackJack) Players() []Player {
	return b.players
}

func (b *blackJack) Winners() []Player {
	return b.winners
}

func (b *blackJack) Ties() []Player {
	return b.ties
}

func (b *blackJack) Losers() []Player {
	return b.losers
}

func (b *blackJack) hostScore() (int, error) {
	var host Player

	for _, player := range b.players {
		if player.Name() == "Host" {
			host = player
		}
	}

	if host != nil {
		return host.Score(), nil
	} else {
		return 0, fmt.Errorf("Host not part of game")
	}
}

func (b *blackJack) sortPlayer(selectPlayer Player) error {

	hostScore, err := b.hostScore()
	if err != nil {
		return err
	}

	isScoreGreaterThanHost := selectPlayer.Score() > hostScore
	isScoreLessThanOrEqualTwentyOne := selectPlayer.Score() <= 21
	isScoreEqualToHost := selectPlayer.Score() == hostScore

	isScoreBetterThanHost := isScoreGreaterThanHost && isScoreLessThanOrEqualTwentyOne

	if isScoreEqualToHost {
		b.ties = append(b.ties, selectPlayer)
	} else if !isScoreBetterThanHost {
		b.losers = append(b.losers, selectPlayer)
	} else if isScoreBetterThanHost {
		b.winners = append(b.winners, selectPlayer)
	}

	return nil
}

func (b *blackJack) removeCardFromDeck(cardQuery CardQuery) (Card, error) {
	newCards := []Card{}
	var foundCard Card

	for _, card := range b.cards {
		if card.Rank == cardQuery.Rank && card.Suit == cardQuery.Suit {
			foundCard = card
		} else {
			newCards = append(newCards, card)
		}
	}

	b.cards = newCards

	if (foundCard != Card{}) {
		return foundCard, nil
	} else {
		return Card{}, fmt.Errorf("Card %v not found", cardQuery)
	}
}

func (b *blackJack) findPlayer(playerName string) (Player, error) {
	var foundPlayer Player

	for _, player := range b.players {
		if player.Name() == playerName {
			foundPlayer = player
		}
	}

	if foundPlayer != nil {
		return foundPlayer, nil
	} else {
		return nil, fmt.Errorf("Can't find %s, %s not part of this game", playerName, playerName)
	}
}

func (b *blackJack) hit(player Player, cardQuery CardQuery) error {
	cardFromDeck, err := b.removeCardFromDeck(cardQuery)
	if err != nil {
		return err
	}

	player.GetCard(cardFromDeck)
	return nil
}

func (b *blackJack) HitPlayer(playerName string, cardQuery CardQuery) {
	player, err := b.findPlayer(playerName)
	if err != nil {
		panic(err)
	}

	err = b.hit(player, cardQuery)
	if err != nil {
		panic(err)
	}
}

func (b *blackJack) RoundComplete() {
	playersExcludingHost := []Player{}

	for _, player := range b.players {
		if player.Name() != "Host" {
			playersExcludingHost = append(playersExcludingHost, player)
		}
	}

	for _, player := range playersExcludingHost {
		b.sortPlayer(player)
	}
}

func (b *blackJack) GetScores() {
	scores := map[string]int{}

	for _, player := range b.players {
		scores[player.Name()] = player.Score()
	}

	fmt.Println("Scores: ", scores)
}
