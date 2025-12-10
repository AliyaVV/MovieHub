package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AliyaVV/MovieHub/internal/repository"
)

func Log_slice(wg *sync.WaitGroup, ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	defer wg.Done()
	prevLenShort := 0
	prevLenExtend := 0
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("logger отменилась")
			return
		default:
			curLenShort := len(repository.SlMovieShort)
			curLenEx := len(repository.SlMovieEx)

			if curLenShort > prevLenShort {
				fmt.Println("больше текущий шорт", curLenShort, "предыдущая", prevLenShort)
				for i := prevLenShort; i < curLenShort; i++ {
					fmt.Println(repository.SlMovieShort[i])
				}
				prevLenShort = curLenShort

			}
			if curLenEx > prevLenExtend {
				fmt.Println("больше текущий лонг")
				for i := prevLenExtend; i < curLenEx; i++ {
					log.Println(repository.SlMovieEx[i])
				}
				prevLenExtend = curLenEx
			}
		}

		fmt.Println("ждем")
	}

}
