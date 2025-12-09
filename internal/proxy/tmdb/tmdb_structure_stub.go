package tmdb

type MovieStubTMDB struct{}

var strSearch = TMDBSearchTitle{
	ID:    50014,
	Title: "The Help",
	Description: `Aibileen Clark is a middle-aged African-American maid who has spent her life raising 
			white children and has recently lost her only son; Minny Jackson is an African-American maid who has often 
			offended her employers despite her family's struggles with money and her desperate need for jobs; 
			and Eugenia \"Skeeter\" Phelan is a young white woman who has recently moved back home after graduating 
			college to find out her childhood maid has mysteriously disappeared. These three stories intertwine to explain 
			how life in Jackson, Mississippi revolves around \"the help\"; yet they are always kept at a certain distance 
			because of racial lines.`,
	Movie_year: "2011-08-09",
	Rating:     8.204,
	Genre:      []int{18},
}

func (*MovieStubTMDB) TMDBSearchTitle(title string) (*RespTMDBSearchTitle, error) {

	return &RespTMDBSearchTitle{
		Results: []TMDBSearchTitle{
			strSearch,
		},
	}, nil
}

func (*MovieStubTMDB) TMDBSearchId(id int) (*RespTMDBMovieDetail, error) {
	return &RespTMDBMovieDetail{
		TMDBSearchTitle: strSearch,
		Budget:          25_000_000,
		Genres: []Genres{
			{
				ID:   18,
				Name: "Drama",
			},
		},
		Country: []string{
			"US",
		},
		DescShort: "Change begins with a whisper",
	}, nil
}
