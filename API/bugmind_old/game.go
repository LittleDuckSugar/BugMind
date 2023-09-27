package bugmind

import (
	"BugMindAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Game = models.GameData{Player: [2]models.Player{}, PlayerTurn: 0, History: make([]models.Play, 0)}

func init() {
	Game.Player[0] = newPlayer("Emma")
	Game.Player[1] = newPlayer("Antoine")

	Game.GeneratePlayersCards(loadDeck())

	Game.History = append(Game.History, models.Play{Command: "init", PlayerId: 255})
}

func InitParty(c *gin.Context) {
	Game.Player[0] = newPlayer("Emma")
	Game.Player[1] = newPlayer("Antoine")

	Game.GeneratePlayersCards(loadDeck())

	Game.History = append(Game.History, models.Play{Command: "init", PlayerId: 255})

	c.JSON(http.StatusOK, Game)
}

func GetParty(c *gin.Context) {
	c.JSON(http.StatusOK, Game)
}

func GetCurrentBoard(c *gin.Context) {
	var board = make(map[string]interface{})

	for _, player := range Game.Player {
		board[player.Name] = player.Board
	}

	c.JSON(http.StatusOK, board)
}

func PostPlay(c *gin.Context) {
	var play models.Play
	c.ShouldBindJSON(&play)

	// Check player turn
	if play.PlayerId != Game.PlayerTurn {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not your turn to play " + Game.Player[play.PlayerId].Name})
		return
	}

	var code int
	var detail string

	// Call method depending on players actions
	switch Game.History[len(Game.History)-1].Command {
	case "init":
		if play.Command == "play" && len(Game.History) == 1 {
			code, detail = Game.Player[play.PlayerId].Play(play.Card)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Command not valid"})
			return
		}
	case "play":
		switch play.Command {
		case "play":
			code, detail = Game.Player[play.PlayerId].Play(play.Card)
		case "attack":
			code, detail = play.Card.Attack(Game.Player[play.PlayerId].Board)
		case "bugmind":
			Game.Player[play.PlayerId].Bugmind(play.Card)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Command not valid"})
			return
		}
	case "attack":
		if play.Command == "defend" {
			Game.Player[play.PlayerId].Defend(play.Card)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Command not valid"})
			return
		}
	case "defend", "mindbug":
		switch play.Command {
		case "play":
			code, detail = Game.Player[play.PlayerId].Play(play.Card)
		case "attack":
			code, detail = play.Card.Attack(Game.Player[play.PlayerId].Board)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Command not valid"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Command not valid"})
		return
	}

	if code%100 != 2 {
		c.JSON(code, gin.H{"error": detail})
		return
	}

	// Save player moves
	Game.History = append(Game.History, play)

	// Set to the other player turn
	if Game.PlayerTurn == 0 {
		Game.PlayerTurn = 1
	} else {
		Game.PlayerTurn = 0
	}

	c.JSON(code, gin.H{"detail": detail})
}
