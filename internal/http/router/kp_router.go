package router

import (
	"github.com/AliyaVV/MovieHub/internal/handler/kphandler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *kphandler.MovieHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/movies/:id", movieHandler.GetMovieById)
	r.GET("/movies/search", movieHandler.GetMovieByTitle)

	return r
}
