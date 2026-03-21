package handler

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

// GetMovieById godoc
// @Summary Поиск фильма по ID
// @Description Возвращает полную информацию о фильме
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} model.Movie_ex
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [get]
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

// GetMovieByTitle godoc
// @Summary Поиск фильма по названию
// @Description Возвращает список фильмов
// @Tags movies
// @Accept json
// @Produce json
// @Param query query string true "Movie title"
// @Success 200 {array} model.Movie_short
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/search [get]
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

// GetMovies godoc
// @Summary Получить список фильмов
// @Description Возвращает все фильмы из БД
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {array} model.Movie_short
// @Failure 500 {object} map[string]string
// @Router /movies/list [get]
func (h *MovieHandler) GetMovies(c *gin.Context) {
	movieList, err := h.service.GetMovies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movieList)
}
