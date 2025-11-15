package model

//структура для агрегированных данных, содержит краткую ифнормацию о фильмах
type Movie_short struct {
	Runame    string
	EnName    string
	MovieType string
	MovieYear int
	Genres    []MovieGenres //берем из MovieEx.go
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
func (m Movie_short) GetKPId() string {
	return m.ExternalId.KPHD
}
