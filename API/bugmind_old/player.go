package bugmind

import "BugMindAPI/models"

// newPlayer generate a new player profil
func newPlayer(name string) models.Player {
	return models.Player{Name: name, HP: 3, BugMind: 2, Deck: make([]models.Card, 5), Hand: make([]models.Card, 5)}
}
