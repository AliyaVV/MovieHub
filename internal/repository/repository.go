package repository

import (
	"context"

	"github.com/AliyaVV/MovieHub/internal/model"
)

type MovieRepository interface {
	SaveMovie(ctx context.Context, movie *model.Movie_ex) (int32, error)
	GetMovieById(ctx context.Context, id int) (*model.Movie_ex, error)
	GetListMovies(ctx context.Context) ([]model.Movie_short, error)
}
