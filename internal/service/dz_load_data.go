package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AliyaVV/MovieHub/internal/repository"
)

func LoadFromFile(filename string, typeslice string) error {
	//data := make([]byte, 1024) //не поняла,как определять длину слайса
	//	_, errRead := io.ReadFull(file, data) // как использовать, если всегда EOF

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("файла нет")
			return err
		}
		fmt.Println("ошибка1", err)
		return err
	}
	if len(data) == 0 {
		fmt.Println("пустой файл")
		return nil
	}

	switch typeslice {
	case "short":
		var tempshort repository.MovieShort
		if err := json.Unmarshal(data, &tempshort); err != nil {
			return err
		}
		repository.SlMovieShort = append(repository.SlMovieShort, tempshort...)

	case "long":
		var templong repository.MovieLong
		if err := json.Unmarshal(data, &templong); err != nil {
			return err
		}
		repository.SlMovieLong = append(repository.SlMovieLong, templong...)
	}
	return nil

}

// не смогла вынести запись в слайсы в универсальную функцию из-за того, что
// не поняла, как объявить temp так, чтобы потом в зависимости от параметра targetslice
// присвоить  ему разные типы
// func WriteSlice(datasl []byte, targetslice string) error {
// 	var temp

// 	if targetslice=="MovieShort"{
// 		var temp repository.MovieShort
// 		if err := json.Unmarshal(datasl, &temp); err != nil {
// 		return err
// 	} else{
// 		var temp repository.MovieLong
// 		if err := json.Unmarshal(datasl, &temp); err != nil {
// 		return err
// 		}
// 	}
// 	}
// 	*sl = append(*sl, temp...)
// 	fmt.Println(*sl)
// }
