package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/internal/service"
)

func main() {
	// for i := 0; i <= 3; i++ {
	// 	service.Structure_Create()
	// 	time.Sleep(2 * time.Second)
	// }
	mainContext, mainCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go service.Log_slice(wg, mainContext)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go service.Structure_Create(wg, mainContext)
		wg.Add(2)
		go repository.Movie_Split(wg, mainContext)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("Сигнал", sig)
	mainCancel()
	wg.Wait()

}
