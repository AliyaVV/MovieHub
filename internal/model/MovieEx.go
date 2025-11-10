package model

type MovieEx struct {
	MovieShort
	Description string
	Top250      int
	Ratings
	Genres   []MovieGenres
	Country  string
	Released string
	Poster   string
	Director []string
	Writer   []string
	Actors   []string
	Awards   []string
	budget   int
}
type MovieGenres struct {
	name string
}
type Ratings struct {
	KP                 float32
	TMDB               float32
	FilmCritic         float32
	filmCritics        float32
	russianFilmCritics float32
}
