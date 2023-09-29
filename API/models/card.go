package models

import (
	"net/http"
	"reflect"
)

type Deck struct {
	Cards [32]Card `json:"cards"`
}

type Card struct {
	Name       string   `json:"name"`
	Power      uint8    `json:"power"`
	Keywords   []string `json:"keywords"`
	Ability    string   `json:"ability"`
	CardNumber uint8    `json:"card-number"`
	Copy       uint8    `json:"copy"`
}

func (card *Card) Attack(playerBoard []Card) (code int, detail string) {
	hasCard := false

	// Check if player have the card on board
	for _, boardCard := range playerBoard {
		if reflect.DeepEqual(boardCard, *card) {
			hasCard = true
			break
		}
	}

	if !hasCard {
		return http.StatusExpectationFailed, "You don't have this card on your board"
	}

	return http.StatusAccepted, card.Name + " is attacking"
}
