package models

type Song struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
	Duration int    `json:"duration"`
}

var Songs = []Song{
	{ID: "1", Title: "Shape of You", Artist: "Ed Sheeran", Genre: "Pop", Year: 2017, Duration: 233},
	{ID: "2", Title: "Bohemian Rhapsody", Artist: "Queen", Genre: "Rock", Year: 1975, Duration: 354},
	{ID: "3", Title: "Blinding Lights", Artist: "The Weeknd", Genre: "Synthpop", Year: 2019, Duration: 200},
	{ID: "4", Title: "Hotel California", Artist: "Eagles", Genre: "Rock", Year: 1976, Duration: 390},
	{ID: "5", Title: "Someone Like You", Artist: "Adele", Genre: "Pop", Year: 2011, Duration: 284},
}

// TODO: Bring your ArtistInfo struct and ValidateSingleAlbum logic from Assignment 2

type ArtistInfo struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
