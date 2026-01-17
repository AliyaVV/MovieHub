package repository

import (
	"context"
	"log"
	"os"

	"fmt"

	"encoding/json"
	"sync"

	"github.com/AliyaVV/MovieHub/internal/model"
)

// Домашка по интерфейсам, в проекте не используется

type MovieModel interface {
	GetKPId() string
}

var Ch = make(chan MovieModel, 3)

// var SlMovieShort = []model.Movie_short{} //делаем публиным
// var SlMovieEx = []model.Movie_ex{}       //делаем публиным
var mtx sync.Mutex //тк происходит запись в слайсы берем обычный mutex, а не RW

type MovieShort []model.Movie_short
type MovieLong []model.Movie_ex

var SlMovieShort MovieShort
var SlMovieLong MovieLong

func (mvsh *MovieShort) Add(data model.Movie_short) {
	mtx.Lock()
	defer mtx.Unlock()
	*mvsh = append(*mvsh, data)

}

func (mvsh MovieShort) AddToFile(filename string) {
	data, err := json.MarshalIndent(mvsh, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, data, 0644)
}

func (mvlng *MovieLong) Add(data model.Movie_ex) {
	mtx.Lock()
	defer mtx.Unlock()
	*mvlng = append(*mvlng, data)

}

func (mvlng MovieLong) AddToFile(filename string) {
	data, err := json.MarshalIndent(mvlng, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, data, 0644)
}

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
			switch elem := chstr.(type) {
			case model.Movie_short:
				SlMovieShort.Add(elem)
				SlMovieShort.AddToFile("shortSlice.json")
				fmt.Println("Movie_short")
			case model.Movie_ex:
				SlMovieLong.Add(elem)
				SlMovieLong.AddToFile("longSlice.json")
				fmt.Println("Movie_ex")
			default:
				fmt.Println("default")
			}
		}

		fmt.Println("split отработала")
	}

}

func Create_file() {
	mtx.Lock()
	defer mtx.Unlock()
	//fileShort, _ := os.Create("shortSlice.json")
	data, err := json.MarshalIndent(SlMovieShort, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("shortSlice.json", data, 0644)

	//fileLong, _ := os.Create("longSlice.json")
	dataLong, err2 := json.MarshalIndent(SlMovieLong, "", "  ")
	if err2 != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("longSlice.json", dataLong, 0644)
}
