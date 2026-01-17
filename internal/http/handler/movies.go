package handler

import (
	"fmt"
	"net/http"

	"github.com/AliyaVV/MovieHub/internal/http/handler/dto"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repository.MovieRepository
}

func New(repo repository.MovieRepository) *Handler {
	return &Handler{repo: repo}
}

func InitHandler(api *gin.Engine, h *Handler) {
	api.GET("/items", h.GetShortMovieList()) //получение списка фильмов
	api.GET("/items/:id", h.GetMovie)        //получение фильма по id
	api.POST("/item", h.AddMovieShort)
	api.PUT("/item/:id", h.EditMovie)
	api.DELETE("/item/:id", h.DeleteMovie)
}

// получение списка фильмов
func (h *Handler) GetShortMovieList() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortMovie := h.repo.GetShort()
		c.JSON(http.StatusOK, shortMovie)
	}
}

// получение фильма по id
func (h *Handler) GetMovie(c *gin.Context) {
	id := c.Param("id")
	onemovie, success := h.repo.GetMovieById(id)
	fmt.Println("succes", success)
	if !success {
		c.JSON(http.StatusNoContent, gin.H{"message": "Movie is not found"})
		return
	}
	c.JSON(http.StatusOK, onemovie)
}

// добавление фильма в слайс и файл
func (h *Handler) AddMovieShort(c *gin.Context) {
	var req dto.AddMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := model.Movie_short{
		Runame:    req.Name,
		MovieType: req.MovieType,
		MovieYear: req.Year,
		ExternalId: model.ExternalId{
			KPHD: req.ID,
		},
	}
	h.repo.Add(movie)
	h.repo.AddToFile("shortSlice.json")
	c.JSON(http.StatusOK, gin.H{
		"message": "movie is added",
	})
}

func (h *Handler) EditMovie(c *gin.Context) {
	id := c.Param("id")
	var req dto.AddMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("id", id)
	if err := h.repo.UpdateMovie(id, req); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "Movie is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "movie is updated",
	})

}

func (h *Handler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteMovieById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "movie is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "movie is deleted"})
}
