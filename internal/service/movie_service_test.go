package service

import (
	"context"
	"errors"
	"testing"

	"github.com/AliyaVV/MovieHub/internal/mapper/tmdbmapper"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
	"github.com/AliyaVV/MovieHub/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockKPResponse() kinopoisk.RespKPSearchID {

	return kinopoisk.RespKPSearchID{
		ID:   101,
		Name: "Movie Test",
		ExternalId: kinopoisk.ExternalId{
			TMDB: 0,
		},
	}
}
func mockKPRespWTMDB() kinopoisk.RespKPSearchID {

	return kinopoisk.RespKPSearchID{
		ID:   101,
		Name: "Movie Test",
		ExternalId: kinopoisk.ExternalId{
			TMDB: 202,
		},
	}
}
func TestGetMovieById(t *testing.T) {
	tests := []struct {
		name       string
		setupMocks func(repo *mocks.MovieRepository, kp *mocks.KPClient, tmdb *mocks.TMDBClient)
		wantErr    bool
	}{
		{
			name: "movie in DB",
			setupMocks: func(repo *mocks.MovieRepository, kp *mocks.KPClient, tmdb *mocks.TMDBClient) {

				repo.On("GetMovieById", mock.Anything, 101).
					Return(&model.Movie_ex{
						Movie_short: model.Movie_short{
							Id:     101,
							Runame: "Movie Test",
						},
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "movie not in DB, movie in KP, TMDB not ",
			setupMocks: func(repo *mocks.MovieRepository, kp *mocks.KPClient,
				tmdb *mocks.TMDBClient) {

				repo.On("GetMovieById", mock.Anything, 101).
					Return(nil, assert.AnError)

				kp.On("GetByID", mock.Anything, 101).
					Return(mockKPResponse(), nil)

				repo.On("SaveMovie",
					mock.Anything,
					mock.MatchedBy(func(m *model.Movie_ex) bool {
						return m.Runame == "Movie Test"
					})).Return(int32(101), nil)
			},
			wantErr: false,
		},
		{
			name: "movie from KP and TMDB",
			setupMocks: func(repo *mocks.MovieRepository, kp *mocks.KPClient,
				tmdb *mocks.TMDBClient) {

				repo.On("GetMovieById", mock.Anything, 101).Return(
					nil, errors.New("GetMovieById: Movie does not exists"))

				kp.On("GetByID", mock.Anything, 101).
					Return(mockKPRespWTMDB(), nil)
				tmdb.On("GetByID", mock.Anything, 202).Return(
					&tmdbmapper.Tmdb_movie_midl{
						ID:    202,
						Title: "Movie Test",
					}, nil,
				)
				repo.On("SaveMovie",
					mock.Anything,
					mock.MatchedBy(func(m *model.Movie_ex) bool {
						return m.Runame == "Movie Test"
					})).Return(int32(101), nil)
			},
			wantErr: false,
		},
		{
			name: "kp error",
			setupMocks: func(repo *mocks.MovieRepository, kp *mocks.KPClient, tmdb *mocks.TMDBClient) {

				repo.On("GetMovieById", mock.Anything, 101).
					Return(nil, assert.AnError)

				kp.On("GetByID", mock.Anything, 101).
					Return(kinopoisk.RespKPSearchID{}, assert.AnError)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			repo := mocks.NewMovieRepository(t)
			kp := mocks.NewKPClient(t)
			tmdb := mocks.NewTMDBClient(t)

			tt.setupMocks(repo, kp, tmdb)

			service := MovieService{
				MovieRepo:     repo,
				KPInterface:   kp,
				TMDBInterface: tmdb,
			}

			_, err := service.GetMovieById(context.Background(), 101)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestGetMovieByTitle(t *testing.T) {
	tests := []struct {
		name       string
		setupMocks func(kp *mocks.KPClient, logger *mocks.SearchLogger)
		wantErr    bool
	}{
		{
			name: "movies found",
			setupMocks: func(kp *mocks.KPClient, logger *mocks.SearchLogger) {

				movies := []*model.Movie_short{
					{
						Id:     101,
						Runame: "Movie Test",
					},
					{
						Id:     102,
						Runame: "Movie Test 2",
					},
				}

				kp.On("SearchByTitle", mock.Anything, "Test Movie").
					Return(movies, nil)

				logger.On("Log", mock.Anything, "Test Movie", 2).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "kp error",
			setupMocks: func(kp *mocks.KPClient, logger *mocks.SearchLogger) {

				kp.On("SearchByTitle", mock.Anything, "Test Movie").
					Return(nil, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "kp ok, logger error",
			setupMocks: func(kp *mocks.KPClient, logger *mocks.SearchLogger) {

				movies := []*model.Movie_short{
					{
						Id:     101,
						Runame: "Movie Test",
					},
				}

				kp.On("SearchByTitle", mock.Anything, "Test Movie").
					Return(movies, nil)

				logger.On("Log", mock.Anything, "Test Movie", 1).
					Return(assert.AnError)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			kp := mocks.NewKPClient(t)
			logger := mocks.NewSearchLogger(t)

			tt.setupMocks(kp, logger)

			service := MovieService{
				KPInterface: kp,
				Logger:      logger,
			}

			res, err := service.GetMovieByTitle(context.Background(), "Test Movie")

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestGetMovies(t *testing.T) {

	tests := []struct {
		name       string
		setupMocks func(repo *mocks.MovieRepository)
		wantErr    bool
	}{
		{
			name: "movies exist",
			setupMocks: func(repo *mocks.MovieRepository) {

				movies := []model.Movie_short{
					{
						Id:     1,
						Runame: "Test Movie",
					},
					{
						Id:     2,
						Runame: "Test Movie 2",
					},
				}

				repo.On("GetListMovies", mock.Anything).
					Return(movies, nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			setupMocks: func(repo *mocks.MovieRepository) {

				repo.On("GetListMovies", mock.Anything).
					Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			repo := mocks.NewMovieRepository(t)

			tt.setupMocks(repo)

			service := MovieService{
				MovieRepo: repo,
			}

			res, err := service.GetMovies(context.Background())

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}
