package tmdbmapper

type tmdb_midl_dto struct {
	ID          int
	Title       string
	Description string
	Movie_year  string
	Rating      float32
	Genre       []int
	Budget      int
	Genres      []Genres
	Country     []string
	DescShort   string
}

type Genres struct {
	ID   int
	Name string
}

//разобраться с актерским составом
