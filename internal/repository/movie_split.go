package repository

import (
	"context"
	"errors"
	"log"
	"os"

	"fmt"

	"encoding/json"
	"sync"

	"github.com/AliyaVV/MovieHub/internal/http/handler/dto"
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

// добавление в слайс списка короткой инф о фильме
func (mvsh *MovieShort) Add(data model.Movie_short) {
	mtx.Lock()
	defer mtx.Unlock()
	*mvsh = append(*mvsh, data)

}

// запись слайса короткой инф о фильме в файл
func (mvsh MovieShort) AddToFile(filename string) error {
	data, err := json.MarshalIndent(mvsh, "", "  ")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	return nil
}

// добавление в слайс детальной инф о фильме
func (mvlng *MovieLong) Add(data model.Movie_ex) {
	mtx.Lock()
	defer mtx.Unlock()
	*mvlng = append(*mvlng, data)

}

// получение слайса короткой инф о фильме
func (mvsh MovieShort) GetShort() []model.Movie_short {
	mtx.Lock()
	defer mtx.Unlock()

	res := make([]model.Movie_short, len(mvsh))
	copy(res, mvsh)

	return res
}

// иннтерфейс для работы с хэндлером
type MovieRepository interface {
	GetShort() []model.Movie_short
	GetMovieById(string) (model.Movie_short, bool)
	Add(model.Movie_short)
	AddToFile(string) error
	UpdateMovie(string, dto.AddMovieRequest) error
	DeleteMovieById(string) error
}

// получение фильма по ид из слайса короткой инф о фильме
func (mvsh MovieShort) GetMovieById(id string) (model.Movie_short, bool) {
	mtx.Lock()
	defer mtx.Unlock()
	for _, i := range mvsh {
		if i.ExternalId.KPHD == id {
			return i, true
		}
	}
	return model.Movie_short{}, false
}

func (mvsh *MovieShort) UpdateMovie(id string, req dto.AddMovieRequest) error {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println("update")
	for i := range *mvsh {
		if (*mvsh)[i].ExternalId.KPHD == id {
			fmt.Println(req.Name)
			(*mvsh)[i].Runame = req.Name
			(*mvsh)[i].MovieType = req.MovieType
			(*mvsh)[i].MovieYear = req.Year
			mvsh.AddToFile("shortSlice.json")
			return nil

		}
	}

	fmt.Println("не найден по ид")
	return errors.New("movie is not found")

}

func (mvsh *MovieShort) DeleteMovieById(id string) error {
	mtx.Lock()
	defer mtx.Unlock()

	index := -1
	for i, movie := range *mvsh {
		if movie.ExternalId.KPHD == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("movie is not found")
	}

	// удаляем
	*mvsh = append((*mvsh)[:index], (*mvsh)[index+1:]...)

	// сохраняем обновлённый слайс в файл
	data, err := json.MarshalIndent(mvsh, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("shortSlice.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// запись слайса детальной инф о фильме в файл
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
