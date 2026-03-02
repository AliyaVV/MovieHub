package tmdbmapper

type Tmdb_movie_midl struct {
	ID          int
	Title       string
	Description string
	Movie_year  string
	Rating      float64
	Budget      int
	Revenue     int
	Genres      []string
	Country     []string
	Slogan      string
}

type Genres struct {
	ID   int
	Name string
}
