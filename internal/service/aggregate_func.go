package service

import (
	"errors"

	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/mapper/tmdbmapper"
	"github.com/AliyaVV/MovieHub/internal/model"
)

func agg_movie_ex(kp *kpmapper.Movie_Entity, tmdb *tmdbmapper.Tmdb_movie_midl) (*model.Movie_ex, error) {
	if kp == nil {
		return nil, errors.New("kp is nil")
	}
	var tmdb_flag bool
	if tmdb != nil && tmdb.ID != 0 && tmdb.Title != "" {
		tmdb_flag = true
	}
	cast := make([]model.Cast, 0, len(kp.Cast))
	for _, val := range kp.Cast {
		if val.Name != "" {
			cast = append(cast, model.Cast{
				Name:        val.Name,
				EnName:      val.EnName,
				Profession:  val.Profession,
				Description: val.Description,
			})
		}
	}
	movie_ex := model.Movie_ex{
		Movie_short: model.Movie_short{
			Id:        kp.ID,
			Runame:    kp.Name,
			EnName:    kp.Name,
			MovieType: kp.MovieType,
			MovieYear: kp.Year,
			Genres:    kp.Genres,
			ExternalId: model.ExternalId{
				TMDB: kp.IDTmdb,
			},
		},
		Source: model.Source{
			KPHD: true,
			TMDB: tmdb_flag,
		},
		Description: kp.Description,
		Top250:      kp.Top250,
		Ratings: model.Ratings{
			KP:                 kp.Ratings.KP,
			FilmCritic:         kp.Ratings.FilmCritics,
			RussianFilmCritics: kp.Ratings.RussianFilmCritics,
		},
		Country: kp.Countries,
		Cast:    cast,
	}
	if tmdb != nil {
		movie_ex.ExternalId.TMDB = tmdb.ID
		movie_ex.Ratings.TMDB = tmdb.Rating
		movie_ex.Budget = tmdb.Budget
		movie_ex.Revenue = tmdb.Revenue
	}

	return &movie_ex, nil
}
