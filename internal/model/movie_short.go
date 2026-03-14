package model

//структура для агрегированных данных, содержит краткую ифнормацию о фильмах
type Movie_short struct {
	Id         int
	Runame     string
	EnName     string
	MovieType  string
	MovieYear  int
	Genres     []string
	ExternalId ExternalId
	Ratings    Ratings
}

type ExternalId struct {
	TMDB int
}

type Source struct {
	TMDB bool
	KPHD bool
}

func NewSource() *Source {
	return &Source{
		TMDB: false,
		KPHD: false,
	}
}
