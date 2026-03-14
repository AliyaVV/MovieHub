package tmdbmapper

import "github.com/AliyaVV/MovieHub/internal/proxy/tmdb"

func MapSearchTitle(resp tmdb.TMDBSearchTitle) *Tmdb_movie_midl {
	return &Tmdb_movie_midl{
		ID:          resp.ID,
		Title:       resp.Title,
		Description: resp.Description,
		Movie_year:  resp.Movie_year,
		Rating:      resp.Rating,
	}
}

func MapSearchDetail(resp tmdb.RespTMDBMovieDetail) *Tmdb_movie_midl {
	genres := make([]string, 0, len(resp.Genres))
	for _, val := range resp.Genres {
		if val.Name != "" {
			genres = append(genres, val.Name)
		}
	}
	return &Tmdb_movie_midl{
		ID:          resp.ID,
		Title:       resp.Title,
		Description: resp.Description,
		Movie_year:  resp.Movie_year,
		Rating:      resp.Rating,
		Budget:      resp.Budget,
		Revenue:     resp.Revenue,
		Genres:      genres,
		Country:     resp.Country,
		Slogan:      resp.Slogan,
	}
}
