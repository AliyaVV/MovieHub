package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AliyaVV/MovieHub/configs"
	"github.com/AliyaVV/MovieHub/internal/external/kinopoiskclient"
	"github.com/AliyaVV/MovieHub/internal/external/tmdbclient"
	"github.com/AliyaVV/MovieHub/internal/handler"
	"github.com/AliyaVV/MovieHub/internal/http/router"
	"github.com/AliyaVV/MovieHub/internal/service"
	"github.com/AliyaVV/MovieHub/internal/storage/postgre"
	"github.com/AliyaVV/MovieHub/internal/storage/redis"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	//"github.com/jackc/pgx"
	"github.com/joho/godotenv"
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

func runMigrations(db *sql.DB) error {
	goose.SetDialect("postgres")

	// Применяем все миграции из папки
	if err := goose.Up(db, "internal/storage/postgre/migrations"); err != nil {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

func main() {
	//ctx := context.Background()

	redisClient := redis.InitRedis()
	defer redisClient.Close()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database")
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	if err = runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	movieRepo := postgre.NewMovieRepository(db)
	logger := redis.NewRedisSearchLogger(redisClient)

	cfg_kp := configs.Config{
		BaseURL: "https://api.poiskkino.dev/v1.4/movie",
		Token:   os.Getenv("TOKEN_KP"),
	}
	cfg_tmdb := configs.Config{
		BaseURL: "https://api.themoviedb.org/3",
		Token:   os.Getenv("TOKEN_TMDB"),
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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// запуск сервера
	go func() {
		log.Println("Server started on :8080")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-sigChan
	log.Println("Shutting down server...")

	// для graceful shutdown
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")

}

/*
func main_for_mongo() {
	//ctx := context.Background()
	mongoClient, _ := mongo.InitMongo()
	mongoDB := mongoClient.Database("movies_db")

	redisClient := redis.InitRedis()

	movieRepo := mongo.NewMongoMovieRepository(mongoDB)
	logger := redis.NewRedisSearchLogger(redisClient)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}
	// urlExample := "postgres://user_ps:user_ps@localhost:5432/movie_db"
	/*conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal("Unable to connect to database")
		fmt.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		log.Fatal("Unable to connect to database")
	}

	cfg_kp := configs.Config{
		BaseURL: "https://api.poiskkino.dev/v1.4/movie",
		Token:   os.Getenv("TOKEN_KP"),
	}
	cfg_tmdb := configs.Config{
		BaseURL: "https://api.themoviedb.org/3",
		Token:   os.Getenv("TOKEN_TMDB"),
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

}*/
