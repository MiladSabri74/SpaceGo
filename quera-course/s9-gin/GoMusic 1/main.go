package main

import (
	"GoMusic/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO: Define GET /songs, GET /songs/:id, and GET /search routes
	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/:id", handlers.GetSongByID)
	r.GET("/search", handlers.SearchSongs)

	r.Run(":8080")
}
