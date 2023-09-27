package models

import (
	"math/rand"
	"time"
)

type GameData struct {
	Player     [2]Player `json:"players"`
	PlayerTurn uint8     `json:"playerTurn"`
	History    []Play    `json:"playHistory"`
}

type Play struct {
	Command  string `json:"command"`
	Card     Card   `json:"card"`
	PlayerId uint8  `json:"playerId"`
}

func (game *GameData) GeneratePlayersCards(deck Deck) {
	rand.NewSource(time.Now().UnixNano())

	cardPos := 0

	for card := 0; card < 20; card++ {
		cardNumber := rand.Intn(32)
		if deck.Cards[cardNumber].Copy != 0 {
			if card < 10 {
				game.Player[card%2].Deck[cardPos] = deck.Cards[cardNumber]
			} else {
				game.Player[card%2].Hand[cardPos-5] = deck.Cards[cardNumber]
			}
		} else {
			card--
			continue
		}
		deck.Cards[cardNumber].Copy--
		if card%2 == 1 {
			cardPos++
		}
	}
}
