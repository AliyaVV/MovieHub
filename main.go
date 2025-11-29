package main

import (
	"sync"

	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/internal/service"
)

func main() {

	// for i := 0; i <= 3; i++ {
	// 	service.Structure_Create()
	// 	time.Sleep(2 * time.Second)
	// }
	wg := &sync.WaitGroup{}

	wg.Add(8)
	for i := 0; i < 2; i++ {
		go repository.Log_slice(wg)
		go service.Structure_Create(wg)
		go repository.Movie_Split(wg)
	}

	wg.Wait()

}
