package service

import (
	"context"

	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
)

type MovieService struct {
	KPInterface kinopoiskclient.KPClient
}

// сервис по получению расширенных данных по ид фильма
func (ms *MovieService) GetMovieById(ctx context.Context, id int) (*kpmapper.Movie_Entity, error) {
	resp, err := ms.KPInterface.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	movie_base := kpmapper.GetBaseMovie(resp)
	kpmapper.MapKPDetailToEntity(movie_base, resp)
	return movie_base, nil
}

func (ms *MovieService) GetMovieByTitle(ctx context.Context, title string) ([]*kpmapper.Movie_Entity, error) {
	resp, err := ms.KPInterface.SearchByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
