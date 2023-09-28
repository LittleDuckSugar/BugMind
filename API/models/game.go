package models

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
