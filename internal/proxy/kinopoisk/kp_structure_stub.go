package kinopoisk

type MovieStubKP struct{}

func (movie *MovieStubKP) KPGetMovieTitle(title string) (*RespKPSearchTitle, error) {
	sSeries := KPSearchTitle{
		ID:          401522,
		Name:        "Как я встретил вашу маму",
		MovieType:   "tv-series",
		Year:        2005,
		Description: "Захватывающий рассказ Теда Мосби своим детям о долгом и необыкновенном пути, который привел его к их матери",
		Genres: []Genre{
			{Name: "комедия"},
			{Name: "мелодрама"},
			{Name: "драма"},
		},
		Ratings: Rating{
			KP:                 8.62,
			Imdb:               8.3,
			FilmCritics:        0,
			RussianFilmCritics: 0,
		},
		Top250: 0,
		Top10:  0,
		Votes: Vote{
			KP:                 261295,
			Imdb:               762000,
			FilmCritics:        0,
			RussianFilmCritics: 2,
		},
	}
	respSeries := RespKPSearchTitle{
		Docs: []KPSearchTitle{sSeries},
	}
	return &respSeries, nil
}

func (mobie *MovieStubKP) KPGetMovieId(id int) (*RespKPSearchID, error) {
	return &RespKPSearchID{
		ID:     261295,
		Slogan: "A love story in reverse",
		Cast: []Actors{
			{
				Name:        "Джош Рэднор",
				EnName:      "Josh Radnor",
				Profession:  "актеры",
				Description: "Ted Mosby",
			},
			{
				Name:        "Нил Патрик Харрис",
				EnName:      "Neil Patrick Harris",
				Profession:  "актеры",
				Description: "Barney Stinson",
			},
			{
				Name:        "Коби Смолдерс",
				EnName:      "Cobie Smulders",
				Profession:  "актеры",
				Description: "Robin Scherbatsky",
			},
			{
				Name:        "Джейсон Сигел",
				EnName:      "Jason Segel",
				Profession:  "актеры",
				Description: "Marshall Eriksen",
			},
			{
				Name:        "Элисон Хэннигэн",
				EnName:      "Alyson Hannigan",
				Profession:  "актеры",
				Description: "Lily Aldrin",
			},
		},
	}, nil
}
