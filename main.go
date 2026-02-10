package main

import (
	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/handler/kphandler"
	"github.com/AliyaVV/MovieHub/internal/http/router"
	"github.com/AliyaVV/MovieHub/internal/service"
)

/*
func old_dz_main{
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
}*/

func main() {
	//ctx := context.Background()
	cfg := kinopoiskclient.Config{
		BaseURL: "https://api.poiskkino.dev/v1.4/movie",
		APIKey:  "APDGYB9-VV5MWKH-K3EZDXM-SQG2YYN",
	}
	kpClient := kinopoiskclient.NewHTTPClient(cfg)
	movieService := &service.MovieService{
		KPInterface: kpClient,
	}
	movieHandler := kphandler.NewMovieHandler(movieService)

	r := router.SetupRouter(movieHandler)

	r.Run(":8080")

}
