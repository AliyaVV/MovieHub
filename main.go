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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/AliyaVV/MovieHub/docs"
)

func runMigrations(db *sql.DB) error {
	goose.SetDialect("postgres")

	if err := goose.Up(db, "internal/storage/postgre/migrations"); err != nil {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

// @title MovieHub API
// @version 1.0
// @description API для работы с фильмами
// @host localhost:8080
// @BasePath /
func main() {

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
