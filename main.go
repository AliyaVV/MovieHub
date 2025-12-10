package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/AliyaVV/MovieHub/internal/service"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// for i := 0; i <= 3; i++ {
	// 	service.Structure_Create()
	// 	time.Sleep(2 * time.Second)
	// }
	mainContext, mainCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(8)

	go service.Log_slice(wg, mainContext)
	for i := 0; i < 2; i++ {
		go service.Structure_Create(wg, mainContext)
		go repository.Movie_Split(wg, mainContext)
	}
	time.Sleep(1 * time.Second)
	mainCancel()
	fmt.Println("cancel")
	wg.Wait()

}
