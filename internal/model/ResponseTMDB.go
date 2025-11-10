package model

//https://api.themoviedb.org/3/search/movie?query=The Departed&include_adult=false&language=en-US&page=1
type RespTMDBSearchTitle struct {
	Results []TMDBSearchTitle
}

type TMDBSearchTitle struct {
	ID          string  `json:"id"`
	Title       string  `json:"original_title"`
	Description string  `json:"overview"`
	Movie_year  string  `json:"release_date"`
	Rating      float32 `json:"vote_average"`
	Genre       []int
}

//https://api.themoviedb.org/3/movie/1422

type RespTMDBMovieDetail struct {
	TMDBSearchTitle
	Budget    int `json:"budget"`
	Genres    []Genres
	Country   string `json:"origin_country"`
	DescShort string `json:"tagline"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*https://api.themoviedb.org/3/movie/{movie_id}/credits
Планирую брать первых 5 актеров
*/
type RespTMDBActors struct {
	ID     string
	Actros []TMDBCast
}

type TMDBCast struct {
	Name      string
	Character string
}
