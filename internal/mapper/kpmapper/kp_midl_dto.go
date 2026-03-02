package kpmapper

type Movie_Entity struct {
	ID          int
	IDTmdb      int
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
	Countries   []string
}

type KPRatings struct {
	KP                 float64
	Imdb               float64
	FilmCritics        float64
	RussianFilmCritics float64
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
