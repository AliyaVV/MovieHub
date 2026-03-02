package repository

import (
	"context"

	"github.com/AliyaVV/MovieHub/internal/model"
)

type MovieRepository interface {
	Upsert(ctx context.Context, movie *model.Movie_short) error
	GetById(ctx context.Context, id int) (*model.Movie_short, error)
}
