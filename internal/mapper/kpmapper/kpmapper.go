package kpmapper

import (
	"errors"

	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

func MapKPTitleToEntity(respTitle kinopoisk.KPSearchTitle) *Movie_Entity {
	genres := make([]string, 0, len(respTitle.Genres))
	for _, val := range respTitle.Genres {
		if val.Name != "" {
			genres = append(genres, val.Name)
		}
	}
	return &Movie_Entity{
		ID:          respTitle.ID,
		Name:        respTitle.Name,
		MovieType:   respTitle.MovieType,
		Year:        respTitle.Year,
		Description: respTitle.Description,
		Genres:      genres,
		Ratings: KPRatings{
			KP:                 respTitle.Ratings.KP,
			Imdb:               respTitle.Ratings.Imdb,
			FilmCritics:        respTitle.Ratings.FilmCritics,
			RussianFilmCritics: respTitle.Ratings.RussianFilmCritics,
		},
		Top250: respTitle.Top250,
		Top10:  respTitle.Top10,
		Votes: Vote{
			KP:                 respTitle.Votes.KP,
			Imdb:               respTitle.Votes.Imdb,
			FilmCritics:        respTitle.Votes.FilmCritics,
			RussianFilmCritics: respTitle.Votes.RussianFilmCritics,
		},
	}
}

// расширение данных о фильме данными ответа по ид
func MapKPDetailToEntity(movie *Movie_Entity, resp kinopoisk.RespKPSearchID) error {
	if resp.ID == 0 {
		return errors.New("empty kinopoisk_id response")
	}
	movie.Slogan = resp.Slogan
	cast := make([]Actors, 0, len(resp.Cast))
	for ind, valc := range resp.Cast {
		if ind >= 6 {
			break
		}
		cast = append(cast, Actors{
			Name:        valc.Name,
			EnName:      valc.EnName,
			Profession:  valc.Profession,
			Description: valc.Description,
		})
	}
	movie.Cast = append(movie.Cast, cast...)

	seasons := make([]Seasons, 0, len(resp.SeasonsInfo))
	for _, vals := range resp.SeasonsInfo {
		seasons = append(seasons, Seasons{
			Number:        vals.Number,
			EpisodesCount: vals.EpisodesCount,
		})
	}
	movie.SeasonsInfo = append(movie.SeasonsInfo, seasons...)

	return nil
}

// маппинг ответа по ид к промежуточной структуре
func GetBaseMovie(resp kinopoisk.RespKPSearchID) *Movie_Entity {
	genres := make([]string, 0, len(resp.Genres))
	for _, val := range resp.Genres {
		if val.Name != "" {
			genres = append(genres, val.Name)
		}
	}
	return &Movie_Entity{
		ID:          resp.ID,
		Name:        resp.Name,
		MovieType:   resp.MovieType,
		Year:        resp.Year,
		Description: resp.Description,
		Genres:      genres,
		Ratings: KPRatings{
			KP:                 resp.Ratings.KP,
			Imdb:               resp.Ratings.Imdb,
			FilmCritics:        resp.Ratings.FilmCritics,
			RussianFilmCritics: resp.Ratings.RussianFilmCritics,
		},
		Votes: Vote{
			KP:                 resp.Votes.KP,
			Imdb:               resp.Votes.Imdb,
			FilmCritics:        resp.Votes.FilmCritics,
			RussianFilmCritics: resp.Votes.RussianFilmCritics,
		},
	}
}
