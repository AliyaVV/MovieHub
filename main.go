package main

import (
	"github.com/AliyaVV/MovieHub/configs"
	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/external/tmdbclient"
	"github.com/AliyaVV/MovieHub/internal/handler"
	"github.com/AliyaVV/MovieHub/internal/http/router"
	"github.com/AliyaVV/MovieHub/internal/service"
	"github.com/AliyaVV/MovieHub/storage/mongo"
	"github.com/AliyaVV/MovieHub/storage/redis"
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
	mongoClient, _ := mongo.InitMongo()
	mongoDB := mongoClient.Database("movies_db")

	redisClient := redis.InitRedis()

	movieRepo := mongo.NewMongoMovieRepository(mongoDB)
	logger := redis.NewRedisSearchLogger(redisClient)

	cfg_kp := configs.Config{
		BaseURL: "https://api.poiskkino.dev/v1.4/movie",
		Token:   "APDGYB9-VV5MWKH-K3EZDXM-SQG2YYN",
	}
	cfg_tmdb := configs.Config{
		BaseURL: "https://api.themoviedb.org/3",
		Token:   `eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzYjNlYTcxMjc4NDVkNDgyY2NhODQxZjU0MjIzNGUxNiIsIm5iZiI6MTc2MjU5OTgwMC43MjIsInN1YiI6IjY5MGYyMzc4ZjFhYWYyZGViMmNlN2ZjMiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.2q38BKb0DVFpMLyUVtTKPDSmplNLVmne0Pi1kPrl-Kg`,
	}
	kpClient := kinopoiskclient.NewHTTPClient(cfg_kp)
	tmdbClient := tmdbclient.NewHTTPClient(cfg_tmdb)
	movieService := &service.MovieService{
		KPInterface:   kpClient,
		TMDBInterface: tmdbClient,
		MovieRepo:     movieRepo,
		Logger:        logger,
	}
	movieHandler := handler.NewMovieHandler(movieService)

	r := router.SetupRouter(movieHandler)

	r.Run(":8080")

}
