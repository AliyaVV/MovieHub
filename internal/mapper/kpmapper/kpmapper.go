package kpmapper

import (
	"errors"
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

func MapKPTitleToEntity(respTitle kinopoisk.KPSearchTitle) (*model.Movie_short, error) {
	genres := make([]string, 0, len(respTitle.Genres))
	for _, val := range respTitle.Genres {
		if val.Name != "" {
			genres = append(genres, val.Name)
		}
	}
	return &model.Movie_short{
		Id:        respTitle.ID,
		Runame:    respTitle.Name,
		EnName:    respTitle.EnName,
		MovieType: respTitle.MovieType,
		MovieYear: respTitle.Year,
		Genres:    genres,
		Ratings: model.Ratings{
			KP:                 respTitle.Ratings.KP,
			FilmCritic:         respTitle.Ratings.FilmCritics,
			RussianFilmCritics: respTitle.Ratings.RussianFilmCritics,
		},
		ExternalId: model.ExternalId{
			TMDB: respTitle.ExternalId.TMDB,
		},
	}, nil
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
	movie.Cast = cast

	seasons := make([]Seasons, 0, len(resp.SeasonsInfo))
	for _, vals := range resp.SeasonsInfo {
		seasons = append(seasons, Seasons{
			Number:        vals.Number,
			EpisodesCount: vals.EpisodesCount,
		})
	}
	movie.SeasonsInfo = seasons

	return nil
}

// маппинг ответа по ид к промежуточной структуре
func GetBaseMovie(resp kinopoisk.RespKPSearchID) (*Movie_Entity, error) {
	if resp.ID == 0 {
		return nil, errors.New("empty kinopoisk_id response")
	}
	genres := make([]string, 0, len(resp.Genres))
	for _, val := range resp.Genres {
		if val.Name != "" {
			genres = append(genres, val.Name)
		}
	}
	countries := make([]string, 0, len(resp.Countries))
	for _, val := range resp.Countries {
		if val.Name != "" {
			countries = append(countries, val.Name)
		}
	}
	fmt.Println("Awards", ConvertAwards(resp.Awards))

	return &Movie_Entity{
		ID:          resp.ID,
		Name:        resp.Name,
		MovieType:   resp.MovieType,
		Year:        resp.Year,
		Description: resp.Description,
		Genres:      genres,
		Ratings: KPRatings{
			KP:                 resp.Ratings.KP,
			FilmCritics:        resp.Ratings.FilmCritics,
			RussianFilmCritics: resp.Ratings.RussianFilmCritics,
		},
		Countries: countries,
		Votes: Vote{
			KP:                 resp.Votes.KP,
			Imdb:               resp.Votes.Imdb,
			FilmCritics:        resp.Votes.FilmCritics,
			RussianFilmCritics: resp.Votes.RussianFilmCritics,
		},
		IDTmdb: resp.ExternalId.TMDB,
		Awards: ConvertAwards(resp.Awards),
	}, nil
}

func ConvertAwards(awards []kinopoisk.KPAward) []string {

	result := make([]string, 0, len(awards))

	for i, a := range awards {
		if i >= 6 {
			break
		}
		title := fmt.Sprintf(
			"%s %d — %s",
			a.Nomination.Award.Title,
			a.Nomination.Award.Year,
			a.Nomination.Title,
		)

		if a.Winning {
			title += " (winner)"
		}

		result = append(result, title)
	}

	return result
}
