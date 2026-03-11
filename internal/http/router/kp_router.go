package router

import (
	"github.com/AliyaVV/MovieHub/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handler.MovieHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/movies/:id", movieHandler.GetMovieById)
	r.GET("/movies/search", movieHandler.GetMovieByTitle)
	r.GET("/movies/list", movieHandler.GetMovies)

	return r
}
