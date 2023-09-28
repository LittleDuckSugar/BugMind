package models

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Rooms contain all rooms
var Rooms = make(map[string]Room)

type Room struct {
	Name         string    `json:"name"`
	PlayerInside uint8     `json:"player-inside"`
	Status       string    `json:"status"`
	Private      bool      `json:"private"`
	MaxPlayer    uint8     `json:"max-player"`
	Players      [2]Player `json:"players"`
	Password     string    `json:"password"`
}

type NewRoom struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	MaxPlayer uint8  `json:"max-player"`
}

type EnterRoom struct {
	Password   string `json:"password"`
	PlayerName string `json:"player-name"`
}

// CheckPassword do what is means, check password
func (room *Room) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(room.Password), []byte(password))
}

// EncryptPassword do what is means, encrypt
func (room *Room) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(room.Password), 14)
	room.Password = string(bytes)
}

// SetUpPlayersGame do what it says, set up players game
func (room *Room) SetUpPlayersGame() {
	rand.NewSource(time.Now().UnixNano())

	var fullDeck Deck

	byteValue, _ := os.ReadFile("./data/fullDeck.json")

	json.Unmarshal(byteValue, &fullDeck)

	players := room.Players

	for playerCount := 0; playerCount < 2; playerCount++ {
		for cardPass := 0; cardPass < 10; cardPass++ {
			cardNumber := rand.Intn(32)
			if fullDeck.Cards[cardNumber].Copy != 0 {
				players[playerCount].Draw = append(players[0].Draw, fullDeck.Cards[cardNumber])
			} else {
				cardPass--
			}
		}
		players[playerCount].DrawCard(5)
		players[playerCount].BugMind = 2
		players[playerCount].HP = 3
	}

	room.Players = players
}
