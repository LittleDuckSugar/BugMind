package models

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
