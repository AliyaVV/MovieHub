package repository

import (
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/model"
)

func Movie_Split(i model.MovieModel) {
	var slMovieShort = []model.Movie_short{}
	var slMovieEx = []model.Movie_ex{}
	//s := []string{}

	switch elem := i.(type) {
	case model.Movie_short:
		slMovieShort = append(slMovieShort, elem)
		fmt.Println("Movie_short")
	case model.Movie_ex:
		slMovieEx = append(slMovieEx, elem)
		fmt.Println("Movie_ex")
	default:
		fmt.Println("default")
	}

}
