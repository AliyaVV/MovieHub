package repository

import (
	"testing"

	"github.com/AliyaVV/MovieHub/internal/model"
)

func TestMovieSplit(t *testing.T) {
	var test_structure = &model.Movie_short{
		Runame:    "Мстители",
		EnName:    "Avengers",
		MovieType: "movie",
	}

	Movie_Split(test_structure)

}
