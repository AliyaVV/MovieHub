package model

// get поиск по названию

type RespKPSearchTitle struct {
	docs []KPSearchTitle
}

type KPSearchTitle struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	MovieType   string  `json:"type"`
	Year        int     `json:"year"`
	Description string  `json:"shortDescription"`
	Genres      []Genre `json:"genres"`
	Poster      string  `json:"poster"`
	Ratings     Rating  `json:"rating"`
	Top250      int     `json:"top250"`
	Top10       int     `json:"top10"`
	Votes       Vote
}
type Genre struct {
	Name string
}
type Rating struct {
	KP                 float32
	Imdb               float32
	Tmdb               float32
	FilmCritics        float32
	RussianFilmCritics float32
}

type Vote struct {
	KP                 int
	Imdb               int
	Tmdb               int
	FilmCritics        int
	RussianFilmCritics int
}

// get https://api.poiskkino.dev/v1.4/movie/263531
type RespKPSearchID struct {
	ID     string
	Slogan string
	Actors
}
type Actors struct {
	Name   string `json:"name"`
	EnName string `json:"enName"`
	Role   string `json:"description"`
}
