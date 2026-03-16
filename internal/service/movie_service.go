package service

import (
	"context"
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/external/tmdbclient"
	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/internal/storage/redis"
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
	//проверяем БД
	movieFromDB, err := ms.MovieRepo.GetMovieById(ctx, id)
	if err == nil {
		fmt.Printf("Movie with ID %d found in DB", id)
		return movieFromDB, nil
	}
	if err.Error() != "GetMovieById: Movie does not exists" {

		fmt.Printf("Movie Found DB error: %v\n", err)
	} else {
		fmt.Printf("Movie with ID %d does not found in DB\n", id)
	}
	resp_kp, err := ms.KPInterface.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	id_tmdb := resp_kp.ExternalId.TMDB
	movie_base, err := kpmapper.GetBaseMovie(resp_kp)
	if err != nil {
		return nil, err
	}
	kpmapper.MapKPDetailToEntity(movie_base, resp_kp)

	if id_tmdb == 0 {
		fmt.Println("Аггрегируем1")
		movie, err = agg_movie_ex(movie_base, nil)
		if err != nil {
			fmt.Println("error aggregate: ", err)
		}
		_, err = ms.MovieRepo.SaveMovie(ctx, movie)
		if err != nil {
			fmt.Println("error SaveMovie: ", err)
		}
		return movie, nil
	} else {

		resp_tmdb, err := ms.TMDBInterface.GetByID(ctx, id_tmdb)
		if err != nil {
			movie, err = agg_movie_ex(movie_base, nil)
			if err != nil {
				fmt.Println("Service GetMovieById:", err)
			}
			_, err = ms.MovieRepo.SaveMovie(ctx, movie)
			if err != nil {
				fmt.Println("error SaveMovie: ", err)
			}
			return movie, nil
		}
		fmt.Println("Аггрегируем3")
		movie, err = agg_movie_ex(movie_base, resp_tmdb)
		_, err = ms.MovieRepo.SaveMovie(ctx, movie)
		if err != nil {
			fmt.Println("error SaveMovie: ", err)
		}

		return movie, nil
	}
}

func (ms *MovieService) GetMovieByTitle(ctx context.Context, title string) ([]*model.Movie_short, error) {
	resp, err := ms.KPInterface.SearchByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	err_log := ms.Logger.Log(ctx, title, len(resp))
	if err_log != nil {
		fmt.Println("Log error: ", err_log)
	}

	return resp, nil
}

func (ms *MovieService) GetMovies(ctx context.Context) ([]model.Movie_short, error) {
	movies, err := ms.MovieRepo.GetListMovies(ctx)
	if err != nil {
		fmt.Println("Service GetMovies error: ", err)
		return nil, err
	}
	return movies, nil
}
