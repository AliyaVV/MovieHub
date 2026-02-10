package kinopoiskclient

import (
	"context"
	"time"

	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

type KPClient interface {
	SearchByTitle(ctx context.Context, title string) ([]*kpmapper.Movie_Entity, error)
	GetByID(ctx context.Context, id int) (kinopoisk.RespKPSearchID, error)
}
