package kinopoiskclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/AliyaVV/MovieHub/configs"
	"github.com/AliyaVV/MovieHub/internal/mapper/kpmapper"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

// реализация KPClient
type client struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

func NewHTTPClient(cfg configs.Config) KPClient {
	return &client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
	}
}

// поиск в кинопоиске по названию фильма
func (cl *client) SearchByTitle(ctx context.Context, title string) ([]*model.Movie_short, error) {
	url := fmt.Sprintf("%s/search?query=%s", cl.baseURL, url.QueryEscape(title))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", cl.token)

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var kpResp kinopoisk.RespKPSearchTitle
	if err := json.NewDecoder(resp.Body).Decode(&kpResp); err != nil {
		return nil, err
	}

	var movies []*model.Movie_short
	for _, doc := range kpResp.Docs {
		movie, err := kpmapper.MapKPTitleToEntity(doc)
		if err != nil {
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// поиск в кинопоиске по ид
func (cl *client) GetByID(ctx context.Context, id int) (kinopoisk.RespKPSearchID, error) {
	url := fmt.Sprintf("%s/%d", cl.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return kinopoisk.RespKPSearchID{}, err
	}
	req.Header.Set("X-API-KEY", cl.token)
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return kinopoisk.RespKPSearchID{}, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return kinopoisk.RespKPSearchID{}, fmt.Errorf(
			"kinopoisk search failed: status=%d body=%s",
			resp.StatusCode,
			string(body),
		)
	}
	defer resp.Body.Close()
	var kpRespId kinopoisk.RespKPSearchID
	if err := json.NewDecoder(resp.Body).Decode(&kpRespId); err != nil {
		return kinopoisk.RespKPSearchID{}, err
	}
	return kpRespId, nil
}
