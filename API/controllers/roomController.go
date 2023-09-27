package controllers

import (
	"BugMindAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoomStatus return the selected room informations
func GetRoomStatusByName(c *gin.Context) {
	if models.Rooms[c.Param("roomName")].Name == "" {
		c.JSON(http.StatusNoContent, gin.H{"error": "There is no room here"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": models.Rooms[c.Param("roomName")]})
}

// GetAllRooms return the all rooms informations
func GetAllRooms(c *gin.Context) {
	totalPrivate, totalPublic := 0, 0

	for _, room := range models.Rooms {
		if room.Type == "private" {
			totalPrivate++
		}

		if room.Type == "public" {
			totalPublic++
		}
	}

	c.JSON(http.StatusOK, gin.H{"rooms": models.Rooms, "private": totalPrivate, "public": totalPublic, "total": len(models.Rooms)})
}

// PostPublicRoom return the selected room informations
func PostPublicRoom(c *gin.Context) {
	var publicRoom models.PublicRoom

	if c.ShouldBindJSON(&publicRoom) != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Error while parsing JSON"})
		return
	}

	// Room already exist ?
	if models.Rooms[publicRoom.Name].Name != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "A room already have this name"})
		return
	}

	// add to memory
	models.Rooms[publicRoom.Name] = models.Room{Name: publicRoom.Name, Status: "created", Type: "public", MaxPlayer: publicRoom.MaxPlayer}

	c.JSON(http.StatusOK, gin.H{"room": models.Rooms[publicRoom.Name]})
}

// PostPrivateRoom return the selected room informations
func PostPrivateRoom(c *gin.Context) {
	var privateRoom models.PrivateRoom

	if c.ShouldBindJSON(&privateRoom) != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Error while parsing JSON"})
		return
	}

	// Room already exist ?
	if models.Rooms[privateRoom.Name].Name != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "A room already have this name"})
		return
	}

	privateRoom.EncryptPassword()

	// add to memory
	models.Rooms[privateRoom.Name] = models.Room{Name: privateRoom.Name, Status: "created", Type: "private", MaxPlayer: privateRoom.MaxPlayer, Password: privateRoom.Password}

	c.JSON(http.StatusOK, gin.H{"room": models.Rooms[privateRoom.Name]})
}
