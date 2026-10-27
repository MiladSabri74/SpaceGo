package main

import (
	"GoMusic/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		public := v1.Group("public")
		{
			public.GET("/songs", handlers.GetSongs)
			public.GET("/songs/:id", handlers.GetSongByID)
			public.GET("/search", handlers.SearchSongs)
		}
		admin := v1.Group("admin")
		{
			admin.POST("/songs", handlers.CreateSong)
			admin.PUT("/songs/:id", handlers.UpdateSong)
		}
	}
	v2 := r.Group("v2")
	{
		v2.GET("songs", handlers.GetSongsV2)
	}

	r.Run(":8080")
}
