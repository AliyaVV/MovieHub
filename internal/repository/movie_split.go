package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/AliyaVV/MovieHub/internal/model"
)

// Домашка по интерфейсам, в проекте не используется

type MovieModel interface {
	GetKPId() string
}

var Ch = make(chan MovieModel, 3)

var SlMovieShort = []model.Movie_short{} //делаем публиным
var SlMovieEx = []model.Movie_ex{}       //делаем публиным
var mtx sync.Mutex                       //тк происходит запись в слайсы берем обычный mutex, а не RW

func Movie_Split(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("split отменилась")
			return
		case chstr, ok := <-Ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("зашли в фор")
			switch elem := chstr.(type) {
			case model.Movie_short:
				mtx.Lock()
				SlMovieShort = append(SlMovieShort, elem)
				fmt.Println("Movie_short")
				mtx.Unlock()
			case model.Movie_ex:
				mtx.Lock()
				SlMovieEx = append(SlMovieEx, elem)
				fmt.Println("Movie_ex")
				mtx.Unlock()
			default:
				fmt.Println("default")
			}
		}

		fmt.Println("split отработала")
	}
}
