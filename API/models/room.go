package models

import "golang.org/x/crypto/bcrypt"

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
