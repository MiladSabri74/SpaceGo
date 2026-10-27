package handlers

import (
	"GoMusic/models"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {
	// TODO: Implement pagination (page, limit) and filtering (genre, artist, year)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}

	genre := c.Query("genre")
	artist := c.Query("artist")
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		year = 0
	}
	var resp = make([]models.Song, 0)

	for _, song := range models.Songs {
		if genre != "" && song.Genre != genre {
			continue
		}
		if artist != "" && song.Artist != artist {
			continue
		}
		if year != 0 && song.Year != year {
			continue
		}
		resp = append(resp, song)
	}
	start := (page - 1) * limit
	end := int(math.Min(
		float64(page*limit),
		float64(len(resp))))
	if start > end {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}
	resp = resp[start:end]
	c.JSON(http.StatusOK, resp)
}

func GetSongByID(c *gin.Context) {
	// TODO: Find song by ID in models.Songs or return 404
	id := c.Param("id")
	for _, song := range models.Songs {
		if id != "" && song.ID == id {
			c.JSON(http.StatusOK, song)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "song not found",
	})
}

func SearchSongs(c *gin.Context) {
	// TODO: Search songs by title or artist matching query param q
	queryField := strings.ToLower(c.Query("q"))
	if queryField == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter 'q' is required",
		})
	}
	var resp = make([]models.Song, 0)

	for _, song := range models.Songs {
		if queryField == strings.ToLower(song.Artist) {
			resp = append(resp, song)
			continue
		}
		if queryField == strings.ToLower(song.Title) {
			resp = append(resp, song)
			continue
		}
	}
	c.JSON(http.StatusOK, resp)
}

func CreateSong(c *gin.Context) {
	// TODO: Bind JSON to models.Song and validate; return 400 with {"error": "validation failed"} on error,
	// or assign ID, append to models.Songs, and return 201 Created with the created song.
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": err.Error(),
		})
		return
	}
	id := len(models.Songs) + 1
	song.ID = strconv.Itoa(id)
	models.Songs = append(models.Songs, song)
	c.JSON(http.StatusOK, song)
}

func UpdateSong(c *gin.Context) {
	// TODO: Retrieve song ID from path param (:id), bind JSON to models.Song and validate,
	// find and update song in models.Songs, and return 200 OK with updated song (or 404 if not found).
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": err.Error(),
		})
		return
	}
	for i, s := range models.Songs {
		if s.ID == song.ID {
			models.Songs[i] = song
			c.JSON(http.StatusOK, song)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Id of song not found in database",
	})
}

func GetSongsV2(c *gin.Context) {
	// TODO: Return all songs in V2 response format: {"count": X, "data": [...]}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}

	genre := c.Query("genre")
	artist := c.Query("artist")
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		year = 0
	}
	var resp = make([]models.Song, 0)

	for _, song := range models.Songs {
		if genre != "" && song.Genre != genre {
			continue
		}
		if artist != "" && song.Artist != artist {
			continue
		}
		if year != 0 && song.Year != year {
			continue
		}
		resp = append(resp, song)
	}
	start := (page - 1) * limit
	end := int(math.Min(
		float64(page*limit),
		float64(len(resp))))
	if start > end {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid paggination parameters",
		})
	}
	resp = resp[start:end]
	c.JSON(http.StatusOK, gin.H{
		"count": len(resp),
		"data":  resp})
}
