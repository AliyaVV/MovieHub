package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/AliyaVV/MovieHub/internal/http/handler"
	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	repository.LoadFromFile("D:\\Projects\\MovieHub\\shortSlice.json", "short")
	repository.LoadFromFile("D:\\Projects\\MovieHub\\longSlice.json", "long")

	mainContext, mainCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go service.Log_slice(wg, mainContext)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go service.Structure_Create(wg, mainContext)
		wg.Add(1)
		go repository.Movie_Split(wg, mainContext)
	}

	r := gin.Default()
	repo := &repository.SlMovieShort
	h := handler.New(repo)
	handler.InitHandler(r, h)
	r.Run(":8080")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("Сигнал", sig)
	mainCancel()
	close(repository.Ch)

	wg.Wait()

}
