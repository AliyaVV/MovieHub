package repository

import (
	"context"

	"github.com/AliyaVV/MovieHub/internal/model"
)

/*
type MovieRepository interface {
	Upsert(ctx context.Context, movie *model.Movie_short) error
	GetById(ctx context.Context, id int) (*model.Movie_short, error)
	GetAll(ctx context.Context) ([]*model.Movie_short, error)   //было для монго
	Create(ctx context.Context, movie *model.Movie_short) error //было для монго
}
*/

type MovieRepository interface {
	SaveMovie(ctx context.Context, movie *model.Movie_ex) (int32, error)
	GetMovieById(ctx context.Context, id int) (*model.Movie_ex, error)
	GetListMovies(ctx context.Context) ([]model.Movie_short, error)
}
