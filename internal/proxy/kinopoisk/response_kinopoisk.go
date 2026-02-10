package kinopoisk

//структура для ответа от Кинопоиска, делаем 2 запроса по фильму и касту фильма
// get поиск по названию

type RespKPSearchTitle struct {
	Docs []KPSearchTitle `json:"docs"`
}

type KPSearchTitle struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	MovieType   string  `json:"type"`
	Year        int     `json:"year"`
	Description string  `json:"description"`
	Genres      []Genre `json:"genres"`
	Ratings     Rating  `json:"rating"`
	Top250      int     `json:"top250"` //только в этом ответе
	Top10       int     `json:"top10"`  //только в этом ответе
	Votes       Vote    `json:"Votes"`
}
type Genre struct {
	Name string `json:"name"`
}
type Rating struct {
	KP                 float32 `json:"kp"`
	Imdb               float32 `json:"imdb"`
	FilmCritics        float32 `json:"filmCritics"`
	RussianFilmCritics float32 `json:"russianFilmCritics"`
}

type Vote struct {
	KP                 int `json:"kp"`
	Imdb               int `json:"imdb"`
	FilmCritics        int `json:"filmCritics"`
	RussianFilmCritics int `json:"russianFilmCritics"`
}

// get https://api.poiskkino.dev/v1.4/movie/263531
type RespKPSearchID struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	MovieType   string    `json:"type"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	Ratings     Rating    `json:"rating"`
	Slogan      string    `json:"slogan"`
	Cast        []Actors  `json:"persons"`
	SeasonsInfo []Seasons `json:"seasonsInfo"`
	Genres      []Genre   `json:"genres"`
	Votes       Vote      `json:"Votes"`
}
type Actors struct {
	Name        string `json:"name"`
	EnName      string `json:"enName"`
	Profession  string `json:"profession"`
	Description string `json:"description"`
}

type Seasons struct {
	Number        int `json:"number"`
	EpisodesCount int `json:"episodesCount"`
}
