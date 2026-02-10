package kphandler

import (
	"net/http"
	"strconv"

	"github.com/AliyaVV/MovieHub/internal/service"
	"github.com/gin-gonic/gin"
)

// структура хэндлеров,одержащая сервис
type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{
		service: s,
	}
}

func (h *MovieHandler) GetMovieById(c *gin.Context) {
	id_str := c.Param("id")
	movie_id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not valid id"})
		return
	}
	movie, err := h.service.GetMovieById(c.Request.Context(), movie_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetMovieByTitle(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter is required",
		})
		return
	}

	movie, err := h.service.GetMovieByTitle(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
