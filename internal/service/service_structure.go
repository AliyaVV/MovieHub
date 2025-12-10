package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/repository"
)

func Structure_Create(wg *sync.WaitGroup, ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("str_create отменилась")
		return
	default:
		short_structure := model.Movie_short{
			Runame:    "Мстители",
			EnName:    "Avengers",
			MovieType: "movie",
			MovieYear: 2012,
			Genres: []model.MovieGenres{
				{
					Name: "фантастика",
				},
				{
					Name: "боевик",
				},
				{
					Name: "фэнтези",
				},
				{
					Name: "приключения",
				},
			},
			ExternalId: model.ExternalId{
				TMDB: "24428",
				KPHD: "263531",
			},
			Source: model.Source{
				TMDB: false,
				KPHD: false,
			},
		}
		ex_structure := model.Movie_ex{
			Movie_short: short_structure,
			Description: `Локи возвращается, и в этот раз он не один. Земля оказывается на грани порабощения, 
		и только лучшие из лучших могут спасти человечество. Глава международной организации Щ. И. Т. Ник Фьюри 
		собирает выдающихся защитников справедливости и добра. 
		Под предводительством Капитана Америки Железный Человек, Тор, 
		Невероятный Халк, Соколиный Глаз и Чёрная Вдова вступают в войну с захватчиком.`,
			Country: "USA",
		}

		repository.Ch <- short_structure
		repository.Ch <- ex_structure
		time.Sleep(2 * time.Second)
		wg.Done()
		fmt.Println("str_create завершилась")
	}

	//repository.Movie_Split(short_structure)
	//repository.Movie_Split(ex_structure)

}
