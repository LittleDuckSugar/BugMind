package controllers

import (
	"BugMindAPI/models"
	"BugMindAPI/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoomStatus return the selected room informations
func GetRoomStatusByName(c *gin.Context) {
	if models.Rooms[c.Param("roomName")].Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no room here"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": models.Rooms[c.Param("roomName")]})
}

// GetAllRooms return the all rooms
func GetAllRooms(c *gin.Context) {
	totalPrivate, totalPublic := 0, 0

	for _, room := range models.Rooms {
		if room.Private {
			totalPrivate++
		} else {
			totalPublic++
		}
	}

	c.JSON(http.StatusOK, gin.H{"rooms": models.Rooms, "private": totalPrivate, "public": totalPublic, "total": len(models.Rooms)})
}

// PostNewRoom create a new room
func PostNewRoom(c *gin.Context) {
	var newRoom models.NewRoom

	if c.ShouldBindJSON(&newRoom) != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Error while parsing JSON"})
		return
	}

	formatedNewRoomName := tools.FormatName(newRoom.Name)

	// check if room already exists
	if models.Rooms[formatedNewRoomName].Name != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "A room already have this name"})
		return
	}

	var room models.Room

	if newRoom.Password != "" {
		room = models.Room{Name: formatedNewRoomName, Status: "created", Private: true, MaxPlayer: newRoom.MaxPlayer, PlayerInside: 0, Password: newRoom.Password}
		room.EncryptPassword()
	} else {
		room = models.Room{Name: formatedNewRoomName, Status: "created", Private: false, MaxPlayer: newRoom.MaxPlayer, PlayerInside: 0}
	}

	// add to memory
	models.Rooms[room.Name] = room

	c.JSON(http.StatusCreated, gin.H{"room": room})
}

// PostPlayerAction player control actions of a user in a room
func PostPlayerAction(c *gin.Context) {

	room := models.Rooms[c.Param("roomName")]

	if room.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no room here"})
		return
	}

	if room.PlayerInside >= room.MaxPlayer {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "No more place in this room"})
		return
	}

	action := c.Param("action")[1:]
	if action == "enter" {
		var enteringRoom models.EnterRoom

		if c.ShouldBindJSON(&enteringRoom) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error while parsing JSON"})
			return
		}

		if room.Private {

			if room.CheckPassword(enteringRoom.Password) != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "wrong password"})
				return
			}
		}

		room.Players[room.PlayerInside] = models.Player{Name: enteringRoom.PlayerName, HP: 3, BugMind: 2, Deck: []models.Card{}, Hand: []models.Card{}, Discard: []models.Card{}, Board: []models.Card{}}

		room.PlayerInside++
	}

	models.Rooms[room.Name] = room

	c.JSON(http.StatusAccepted, gin.H{"room-selected": room, "action": action})
}
