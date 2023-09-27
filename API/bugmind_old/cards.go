package bugmind

import (
	"BugMindAPI/models"
	"encoding/json"
	"os"
)

// loadDeck returns the entire deck of the game
func loadDeck() models.Deck {
	var fullDeck models.Deck

	byteValue, _ := os.ReadFile("./bugmind/data/fullDeck.json")

	json.Unmarshal(byteValue, &fullDeck)

	return fullDeck
}
