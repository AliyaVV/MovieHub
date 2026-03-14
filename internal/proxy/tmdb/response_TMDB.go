package tmdb

//структуры для ответоа от TMDB
//https://api.themoviedb.org/3/search/movie?query=The Departed&include_adult=false&language=en-US&page=1
type RespTMDBSearchTitle struct {
	Results []TMDBSearchTitle `json:"results"`
}

type TMDBSearchTitle struct {
	ID          int     `json:"id"`
	Title       string  `json:"original_title"`
	Description string  `json:"overview"`
	Movie_year  string  `json:"release_date"`
	Rating      float64 `json:"vote_average"`
}

//https://api.themoviedb.org/3/movie/1422

type RespTMDBMovieDetail struct {
	ID          int      `json:"id"`
	Title       string   `json:"original_title"`
	Description string   `json:"overview"`
	Movie_year  string   `json:"release_date"`
	Rating      float64  `json:"vote_average"`
	Budget      int      `json:"budget"`
	Genres      []Genres `json:"genres"`
	Country     []string `json:"origin_country"`
	Revenue     int      `json:"revenue"`
	Slogan      string   `json:"tagline"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*https://api.themoviedb.org/3/movie/{movie_id}/credits
Планирую брать первых 6 актеров
*/
type RespTMDBActors struct {
	ID     string     `json:"id"`
	Actors []TMDBCast `json:"cast"`
}

type TMDBCast struct {
	Name      string `json:"name"`
	Character string `json:"character"`
	Type      string `json:"known_for_department"`
}
