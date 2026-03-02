package service

import (
	"context"
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/external/tmdbclient"
	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/storage/redis"
)

type MovieService struct {
	KPInterface   kinopoiskclient.KPClient
	TMDBInterface tmdbclient.TMDBClient

	MovieRepo repository.MovieRepository
	Logger    redis.SearchLogger
}

// сервис по получению расширенных данных по ид фильма
func (ms *MovieService) GetMovieById(ctx context.Context, id int) (*model.Movie_ex, error) {
	var movie *model.Movie_ex
	resp_kp, err := ms.KPInterface.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	movie_base, err := kpmapper.GetBaseMovie(resp_kp)
	if err != nil {
		return nil, err
	}
	kpmapper.MapKPDetailToEntity(movie_base, resp_kp)
	id_tmdb := movie_base.IDTmdb
	if id_tmdb == 0 {
		movie, err = agg_movie_ex(movie_base, nil)
	} else {

		resp_tmdb, err := ms.TMDBInterface.GetByID(ctx, id_tmdb)
		if err != nil {
			movie, err = agg_movie_ex(movie_base, nil)
		}

		movie, err = agg_movie_ex(movie_base, resp_tmdb)
	}

	return movie, nil
}

func (ms *MovieService) GetMovieByTitle(ctx context.Context, title string) ([]*model.Movie_short, error) {
	resp, err := ms.KPInterface.SearchByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	for _, movie := range resp {
		fmt.Println(movie)
		if err := ms.MovieRepo.Upsert(ctx, movie); err != nil {
			fmt.Println("Ошибка апсерта в Монго", err)
			continue
		}
	}
	err_log := ms.Logger.Log(ctx, title, len(resp))
	if err_log != nil {
		fmt.Println("Ошибка логирования", err_log)
	}

	return resp, nil
}
