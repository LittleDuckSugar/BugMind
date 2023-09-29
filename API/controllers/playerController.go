package controllers

import (
	"BugMindAPI/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPlayerInfo return every game of a player
func GetPlayerInfo(c *gin.Context) {
	playerName := c.Param("playerName")

	playerInfos := make(map[string]models.Player)
	var currentGames uint8

	for _, room := range models.Rooms {
		for _, player := range room.Players {
			if player.Name == playerName {
				currentGames++
				playerInfos[room.Name] = player
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"current-games": currentGames, "player-infos": playerInfos})
}

// PostPlayerActionsInRoom control actions for a room (create, enter, quit, delete)
func PostPlayerActionsInRoom(c *gin.Context) {
	action := c.Param("action")[1:]
	roomName := c.Param("roomName")
	playerName := c.Param("playerName")

	var playerPos int = 404

	for _, room := range models.Rooms {
		if room.Name == roomName {
			for pos, player := range room.Players {
				if player.Name == playerName {
					playerPos = pos
					break
				}
			}
		}

		if playerPos != 404 {
			break
		}
	}

	// player exist in the choosen room
	if playerPos != 404 {
		if playerPos == models.Rooms[roomName].PlayerTurn {
			switch action {
			case "play":
				// use card from POST
				// models.Rooms[roomName].Players[playerPos].Play(models.Card{})
				fmt.Println("playing")
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not your turn to play"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player doesn't exist in this room or the room doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": playerPos != 404})
}

// GetAllPlayers return every player
func GetAllPlayers(c *gin.Context) {

	playersSaw := make([]string, 0)

	for _, room := range models.Rooms {
		for _, player := range room.Players {
			if !isInside(player, playersSaw) && player.Name != "" {
				playersSaw = append(playersSaw, player.Name)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"total-players": len(playersSaw), "players-name": playersSaw})
}

func isInside(player models.Player, list []string) bool {
	var inside bool = false

	for _, element := range list {
		if element == player.Name {
			inside = true
		}
	}

	return inside
}
