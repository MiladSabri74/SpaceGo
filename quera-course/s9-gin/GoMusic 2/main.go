package main

import (
	"GoMusic/handlers"
	"GoMusic/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(models.ValidateSingleAlbum, models.Song{})
	}

	// TODO: Define GET /songs, GET /songs/:id, and GET /search routes
	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/:id", handlers.GetSongByID)
	r.GET("/search", handlers.SearchSongs)
	// TODO: Register custom structural validation (ValidateSingleAlbum)
	// TODO: Register POST /songs and PUT /songs/:id routes
	r.POST("/songs", handlers.CreateSong)
	r.PUT("/songs/:id", handlers.UpdateSong)
	r.Run(":8080")
}
