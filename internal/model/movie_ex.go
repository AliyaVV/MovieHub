package model

//Структура для агрегированных данных из двух источников, содержит полные расширенные данные
type Movie_ex struct {
	Movie_short
	Description string
	Top250      int
	Ratings
	Country []string
	Poster  string
	Cast    []Cast
	Awards  []string
	Budget  int
	Revenue int
	Seasons []Seasons
	Source  Source
}
type MovieGenres struct {
	Name string
}
type Ratings struct {
	KP                 float64
	TMDB               float64
	FilmCritic         float64
	RussianFilmCritics float64
}

type Cast struct {
	Name        string
	EnName      string
	Profession  string
	Description string
}

type Seasons struct {
	Number        int
	EpisodesCount int
}
