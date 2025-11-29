package service

import (
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/proxy/kinopoisk"
)

func CreateMovieShort(title string) {
	strSearch := kinopoisk.MovieStubKP{}
	point, _ := strSearch.KPGetMovieTitle(title)
	fmt.Println(point)
}
