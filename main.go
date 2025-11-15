package main

import (
	"time"

	"github.com/AliyaVV/MovieHub/internal/service"
)

func main() {

	for i := 0; i <= 3; i++ {
		service.Structure_Create()
		time.Sleep(5 * time.Second)
		i++
	}

}
