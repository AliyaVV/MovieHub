package kpmapper

type Movie_Entity struct {
	ID          int
	Name        string
	MovieType   string
	Year        int
	Description string
	Genres      []string
	Ratings     KPRatings
	Top250      int
	Top10       int
	Votes       Vote
	Slogan      string
	Cast        []Actors
	SeasonsInfo []Seasons
}

type KPRatings struct {
	KP                 float32
	Imdb               float32
	FilmCritics        float32
	RussianFilmCritics float32
}

type Vote struct {
	KP                 int
	Imdb               int
	FilmCritics        int
	RussianFilmCritics int
}

type Actors struct {
	Name        string
	EnName      string
	Profession  string
	Description string
}

type Seasons struct {
	Number        int
	EpisodesCount int
}
