package main

import (
	"BugMindAPI/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Tells to gin if we are in a dev environment or not
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// Tells to gin to force color in shell
	gin.ForceConsoleColor()

	router := gin.Default()

	router.Use(cors.Default())

	// bugmindRoute := router.Group("/bugmind")
	// {
	// 	bugmindRoute.GET("/party/0", bugmind.GetParty)
	// }

	gamePath := router.Group("/bugmind")
	{
		roomPath := gamePath.Group("/room")
		{
			roomPath.GET("/:roomName", controllers.GetRoomStatusByName)
			roomPath.POST("/new-public-room", controllers.PostPublicRoom)
			roomPath.POST("/new-private-room", controllers.PostPrivateRoom)
		}
		gamePath.GET("/rooms", controllers.GetAllRooms)
	}

	// router.GET("/bugmind/party/0", bugmind.GetParty)
	// router.GET("/bugmind/party/0/start", bugmind.InitParty)
	// router.GET("/bugmind/party/0/ingame", bugmind.GetCurrentBoard)
	// router.POST("/bugmind/room/:roomName/play", bugmind.PostPlay)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
