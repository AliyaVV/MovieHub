package dto

type AddMovieRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	MovieType   string `json:"type" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Description string `json:"shortDescription"`
	Top250      int    `json:"top250"`
	Top10       int    `json:"top10"`
}
