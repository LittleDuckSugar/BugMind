package models

import (
	"fmt"
	"net/http"
	"reflect"
)

type Player struct {
	Name    string `json:"name"`
	HP      uint8  `json:"hp"`
	BugMind uint8  `json:"bugmindLeft"`
	Deck    []Card `json:"deck"`
	Hand    []Card `json:"hand"`
	Discard []Card `json:"discard"`
	Board   []Card `json:"board"`
}

func (player *Player) Bugmind(card Card) {
	fmt.Println(player.Name, "is Bugminding", card.Name)
}

// Play put the selected card on the board and
func (player *Player) Play(card Card) (code int, detail string) {
	hasCard := false
	cardPos := 0

	// Check if player have the card he want to play
	for index, handCard := range player.Hand {
		if reflect.DeepEqual(handCard, card) {
			hasCard = true
			cardPos = index
			break
		}
	}

	if !hasCard {
		return http.StatusBadRequest, "You don't have this card in your hand"
	}

	// putting the card on board
	player.Board = append(player.Board, card)

	// removing card from hand
	player.Hand = append(player.Hand[:cardPos], player.Hand[cardPos+1:]...)

	// draw a card from deck if available
	if len(player.Deck) != 0 {
		player.Hand = append(player.Hand, player.Deck[0])
		player.Deck = player.Deck[1:]
	}

	return http.StatusAccepted, card.Name + " played"
}

func (player *Player) Defend(card Card) {

	// check if the defender have cards on board
	if len(player.Board) == 0 {
		player.HP--
	}

	if card.Name == "body" {
		player.HP--
	}

	fmt.Println(player.Name, "is defending")
}

func (player *Player) isAlive() bool {
	return player.HP <= 0
}
