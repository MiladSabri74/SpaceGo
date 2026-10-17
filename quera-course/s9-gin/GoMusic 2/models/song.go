package models

import (
	"github.com/go-playground/validator/v10"
)

// TODO: Define ArtistInfo struct with fields:
// - Name (string, required)
// - Email (string, required, email format)
type ArtistInfo struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// TODO: Add required validation tags to Song struct according to problem description:
// - Title: required
// - Genre: required, oneof=Pop Rock Jazz
// - Year: required, lte=2026
// - Duration: required, gt=0
// - IsSingle: bool
// - AlbumName: string
type Song struct {
	ID         string     `json:"id"`
	Title      string     `json:"title" binding:"required"`
	ArtistInfo ArtistInfo `json:"artist_info"`
	Genre      string     `json:"genre" binding:"required,oneof=Pop Rock Jazz"`
	Year       int        `json:"year" binding:"required,lte=2026"`
	Duration   int        `json:"duration" binding:"required,gt=0"`
	IsSingle   bool       `json:"is_single"`
	AlbumName  string     `json:"album_name"`
}

// Songs is an in-memory slice simulating a database
var Songs = []Song{
	{
		ID:    "1",
		Title: "Shape of You",
		ArtistInfo: ArtistInfo{
			Name:  "Ed Sheeran",
			Email: "ed@sheeran.com",
		},
		Genre:     "Pop",
		Year:      2017,
		Duration:  233,
		IsSingle:  false,
		AlbumName: "Divide",
	},
}

// TODO: Implement custom validation ValidateSingleAlbum
// It must ensure that if IsSingle is true, AlbumName must be empty
func ValidateSingleAlbum(sl validator.StructLevel) {
	song := sl.Current().Interface().(Song)

	if song.IsSingle == true && song.AlbumName != "" {
		sl.ReportError(song.AlbumName, "AlbumName", "album_name", "single_without_album_name", "")
	}
}
