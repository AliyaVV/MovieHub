package model

//структура для агрегированных данных, содержит краткую ифнормацию о фильмах
type MovieShort struct {
	runame    string
	enName    string
	movieType string
	movieYear string
	genres    []Genre //берем из MovieEx.go
	ExternalId
	Source
}

type ExternalId struct {
	TMDB string
	KPHD string
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
