package kinopoiskclient

import (
	"context"

	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

// type Config struct {
// 	BaseURL string
// 	APIKey  string
// 	Timeout time.Duration
// }

type KPClient interface {
	SearchByTitle(ctx context.Context, title string) ([]*model.Movie_short, error)
	GetByID(ctx context.Context, id int) (kinopoisk.RespKPSearchID, error)
}
