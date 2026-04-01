package tmdbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/AliyaVV/MovieHub/configs"
	"github.com/AliyaVV/MovieHub/internal/mapper/tmdbmapper"
	"github.com/AliyaVV/MovieHub/internal/proxy/tmdb"
)

type TMDBClient interface {
	SearchByTitle(ctx context.Context, title string) ([]*tmdbmapper.Tmdb_movie_midl, error)
	GetByID(ctx context.Context, id int) (*tmdbmapper.Tmdb_movie_midl, error)
}

type client struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

func NewHTTPClient(cfg configs.Config) TMDBClient {
	return &client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
	}
}

func (cl *client) SearchByTitle(ctx context.Context, title string) ([]*tmdbmapper.Tmdb_movie_midl, error) {
	url := fmt.Sprintf("%s/search/movie?query=%s", cl.baseURL, url.QueryEscape(title))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+cl.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		fmt.Println("err1", err)
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error close body", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"tmdb search title failed: status=%d body=%s",
			resp.StatusCode,
			string(body),
		)
	}
	var tmdbResp tmdb.RespTMDBSearchTitle
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResp); err != nil {
		return nil, err
	}
	var movies []*tmdbmapper.Tmdb_movie_midl
	for _, res := range tmdbResp.Results {
		movie := tmdbmapper.MapSearchTitle(res)
		movies = append(movies, movie)
	}
	return movies, nil
}

func (cl *client) GetByID(ctx context.Context, id int) (*tmdbmapper.Tmdb_movie_midl, error) {
	url := fmt.Sprintf("%s/movie/%d", cl.baseURL, id)
	fmt.Println("url", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+cl.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error close body", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"tmdb search by id failed: status=%d body=%s",
			resp.StatusCode,
			string(body),
		)
	}
	fmt.Println(resp.Body)
	var tmdbResp tmdb.RespTMDBMovieDetail
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResp); err != nil {
		return nil, err
	}
	movie := tmdbmapper.MapSearchDetail(tmdbResp)
	return movie, nil

}
