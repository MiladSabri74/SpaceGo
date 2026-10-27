package main

import (
	"GoMusic/handlers"
	"GoMusic/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("api")
	api.Use(middlewares.RateLimiter(), middlewares.CustomRecovery())

	v1 := api.Group("v1")
	{
		public := v1.Group("public")
		{
			public.GET("/songs", handlers.GetSongs)
			public.GET("/songs/:id", handlers.GetSongByID)
			public.GET("/search", handlers.SearchSongs)
		}
		admin := v1.Group("admin")
		admin.Use(gin.BasicAuth(gin.Accounts{"admin": "secret"}))
		{
			admin.POST("/songs", handlers.CreateSong)
			admin.PUT("/songs/:id", handlers.UpdateSong)
		}
	}
	v2 := api.Group("v2")
	{
		v2.GET("songs", handlers.GetSongsV2)
	}

	r.Run(":8080")
}
