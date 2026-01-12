package repository

import (
	"encoding/json"
	"fmt"
	"os"
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
		var tempshort MovieShort
		if err := json.Unmarshal(data, &tempshort); err != nil {
			return err
		}
		SlMovieShort = append(SlMovieShort, tempshort...)

	case "long":
		var templong MovieLong
		if err := json.Unmarshal(data, &templong); err != nil {
			return err
		}
		SlMovieLong = append(SlMovieLong, templong...)
	}
	return nil

}

// не смогла вынести запись в слайсы в универсальную функцию из-за того, что
// не поняла, как объявить temp так, чтобы потом в зависимости от параметра targetslice
// присвоить  ему разные типы
// func WriteSlice(datasl []byte, targetslice string) error {
// 	var temp

// 	if targetslice=="MovieShort"{
// 		var temp MovieShort
// 		if err := json.Unmarshal(datasl, &temp); err != nil {
// 		return err
// 	} else{
// 		var temp MovieLong
// 		if err := json.Unmarshal(datasl, &temp); err != nil {
// 		return err
// 		}
// 	}
// 	}
// 	*sl = append(*sl, temp...)
// 	fmt.Println(*sl)
// }
