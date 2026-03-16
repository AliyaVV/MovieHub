package postgre

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/model"
	repository "github.com/AliyaVV/MovieHub/internal/storage/postgre/sqlc"
)

type MovieRepo struct {
	mvdb *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepo {
	return &MovieRepo{
		mvdb: db,
	}
}

func NullToString(param string) sql.NullString {
	return sql.NullString{
		String: param,
		Valid:  param != "",
	}
}
func NullToInt(numb int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: numb,
		Valid: numb != 0,
	}
}
func NullToInt32(numb int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: numb,
		Valid: numb != 0,
	}
}
func NullToFloat64(fl float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: fl,
		Valid:   fl != 0,
	}
}

// метод SaveMovie
func (mv *MovieRepo) SaveMovie(ctx context.Context, movie *model.Movie_ex) (int32, error) {

	trx, err := mv.mvdb.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Transaction started: %v\n", trx)
	q := repository.New(trx)

	defer trx.Rollback()
	fmt.Println("kp_id", movie.Id)
	fmt.Println("TMDB", movie.ExternalId.TMDB)

	movieID, err := q.SaveMovie(ctx, repository.SaveMovieParams{
		RuName:      NullToString(movie.Runame),
		EnName:      NullToString(movie.EnName),
		Year:        NullToInt32(int32(movie.MovieYear)),
		KpID:        NullToInt32(int32(movie.Id)),
		MovieType:   NullToString(movie.MovieType),
		Description: NullToString(movie.Description),
		Top250:      NullToInt32(int32(movie.Top250)),
		Budget:      NullToInt32(int32(movie.Budget)),
		Revenue:     NullToInt32(int32(movie.Revenue)),
		TmdbID:      NullToInt32(int32(movie.ExternalId.TMDB)),
	})
	if err != nil {
		fmt.Println("err1", err)
		return 0, err
	}
	fmt.Println("TMDB rating", movie.Movie_short.Ratings.TMDB)
	fmt.Println("KP rating", movie.Movie_short.Ratings.KP)
	fmt.Println("FilmCritic rating", movie.Movie_short.Ratings.FilmCritic)

	err = q.SaveRating(ctx, repository.SaveRatingParams{
		MovieID:            NullToInt32(movieID),
		Kp:                 NullToFloat64(movie.Movie_short.Ratings.KP),
		Tmdb:               NullToFloat64(movie.Movie_short.Ratings.TMDB),
		FilmCritic:         NullToFloat64(movie.Movie_short.Ratings.FilmCritic),
		RussianFilmCritics: NullToFloat64(movie.Movie_short.Ratings.RussianFilmCritics),
	})
	if err != nil {
		return 0, err
	}

	for _, genre := range movie.Genres {
		genre_id, err := q.GetGenreByName(ctx, NullToString(genre))
		if err != nil {
			return 0, err
		}
		err = q.SaveGenre(ctx, repository.SaveGenreParams{
			MovieID: movieID,
			GenreID: genre_id,
		})

		if err != nil {
			return 0, err
		}
	}

	for _, cast := range movie.Cast {

		err = q.SaveCast(ctx, repository.SaveCastParams{
			MovieID:     NullToInt32(movieID),
			Name:        NullToString(cast.Name),
			EnName:      NullToString(cast.EnName),
			Profession:  NullToString(cast.Profession),
			Description: NullToString(cast.Description),
		})

		if err != nil {
			return 0, err
		}
	}
	if err = trx.Commit(); err != nil {
		fmt.Println("Commit Error", err)
		return 0, err
	}
	return movieID, nil

}

func ConvertGenres(genr []repository.Genre) []string {
	genres := make([]string, 0, len(genr))
	for _, g := range genr {
		genres = append(genres, g.Name.String)
	}
	return genres
}

func convertCast(cast []repository.Cast) []model.Cast {
	result := make([]model.Cast, len(cast))
	for i, c := range cast {
		result[i] = model.Cast{
			Name:        c.Name.String,
			EnName:      c.EnName.String,
			Profession:  c.Profession.String,
			Description: c.Description.String,
		}
	}
	return result
}

// метод GetMovieById
func (mv *MovieRepo) GetMovieById(ctx context.Context, id int) (*model.Movie_ex, error) {
	fmt.Printf("GetMovieById called with id: %d\n", id)
	q := repository.New(mv.mvdb)

	movie, err := q.GetMovie(ctx, NullToInt32(int32(id)))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("GetMovieById: Movie does not exists")
	}
	if err != nil {
		return nil, fmt.Errorf("GetMovieById: error to get movie: %w", err)
	}

	genres, err := q.GetGenres(ctx, int32(id))
	if err != nil {
		return nil, fmt.Errorf("GetMovieById: error to get genres: %w", err)
	}
	cast, err := q.GetCast(ctx, NullToInt32(int32(id)))
	if err != nil {
		return nil, fmt.Errorf("GetMovieById: error to get cast: %w", err)
	}
	return &model.Movie_ex{
		Movie_short: model.Movie_short{
			Id:        int(movie.KpID.Int32),
			Runame:    movie.RuName.String,
			EnName:    movie.EnName.String,
			MovieType: movie.MovieType.String,
			MovieYear: int(movie.Year.Int32),
			Genres:    ConvertGenres(genres),
			ExternalId: model.ExternalId{
				TMDB: int(movie.TmdbID.Int32),
			},
			Ratings: model.Ratings{
				KP:                 movie.Kp.Float64,
				TMDB:               movie.Tmdb.Float64,
				FilmCritic:         movie.FilmCritic.Float64,
				RussianFilmCritics: movie.RussianFilmCritics.Float64,
			},
		},
		Top250:      int(movie.Top250.Int32),
		Description: movie.Description.String,
		Budget:      int(movie.Budget.Int32),
		Revenue:     int(movie.Revenue.Int32),
		Cast:        convertCast(cast),
	}, nil
}

func (mv *MovieRepo) GetListMovies(ctx context.Context) ([]model.Movie_short, error) {

	q := repository.New(mv.mvdb)

	basicMovies, err := q.GetListMovies(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetListMovies:error to get movies list: %w", err)
	}

	movies := make([]model.Movie_short, 0, len(basicMovies))

	for _, basic := range basicMovies {
		genres, err := q.GetGenres(ctx, basic.ID)
		if err != nil {
			fmt.Println("GetListMovies:error get genres:", basic.ID, err)
		}
		fmt.Println("genres ", genres)
		genres_list := ConvertGenres(genres)
		fmt.Println("genres ", genres_list)
		movieShort := model.Movie_short{
			Id:        int(basic.KpID.Int32),
			Runame:    basic.RuName.String,
			EnName:    basic.EnName.String,
			MovieType: basic.MovieType.String,
			MovieYear: int(basic.Year.Int32),
			ExternalId: model.ExternalId{
				TMDB: int(basic.TmdbID.Int32),
			},
			Ratings: model.Ratings{
				KP:   basic.Kp.Float64,
				TMDB: basic.Tmdb.Float64,
			},
			Genres: genres_list,
		}

		movies = append(movies, movieShort)
	}

	return movies, nil
}

// Добавь этот метод для тестирования
func (mv *MovieRepo) TestDirectInsert(ctx context.Context) error {
	// Прямой INSERT без транзакции
	q := repository.New(mv.mvdb)

	id, err := q.SaveMovie(ctx, repository.SaveMovieParams{
		RuName: NullToString("Тестовый фильм"),
		EnName: NullToString("Test Movie"),
		Year:   NullToInt32(2026),
	})

	if err != nil {
		return fmt.Errorf("direct insert failed: %w", err)
	}

	fmt.Printf("Direct insert successful, got ID: %d\n", id)
	return nil
}
