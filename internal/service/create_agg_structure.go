package service

import (
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
	"github.com/AliyaVV/MovieHub/internal/proxy/tmdb"
)

func CreateMovieShort(title string) {
	searchKP := kinopoisk.MovieStubKP{}
	stubKPTitle, _ := searchKP.KPGetMovieTitle(title)
	fmt.Println(stubKPTitle)
	fmt.Println("------------")
	searchTMDB := tmdb.MovieStubTMDB{}
	stubTMDBTitle, _ := searchTMDB.TMDBSearchTitle(title)
	fmt.Println(stubTMDBTitle)

	// testMovieShort := model.Movie_short{
	// 	Runame: stubKPTitle.Docs[0].Name,
	// }

}
