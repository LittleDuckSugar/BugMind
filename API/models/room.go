package models

import "golang.org/x/crypto/bcrypt"

var Rooms = make(map[string]Room)

type Room struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	MaxPlayer int       `json:"max-player"`
	Players   [2]Player `json:"players"`
	Password  string    `json:"password"`
}

type PublicRoom struct {
	Name      string `json:"name"`
	MaxPlayer int    `json:"max-player"`
}

type PrivateRoom struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	MaxPlayer int    `json:"max-player"`
}

func (privateRoom *PrivateRoom) EncryptPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(privateRoom.Password), 14)
	privateRoom.Password = string(bytes)
}
