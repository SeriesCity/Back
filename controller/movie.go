package controller

import (
	"github.com/SeriesCity/Gateway/internal/core/ports"
	"github.com/SeriesCity/Gateway/internal/core/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MovieHandler struct {
	svc ports.MovieHandlerContract
}

func NewMovieHandler() *MovieHandler {
	svc := services.NewMovieService()
	return &MovieHandler{
		svc: svc,
	}
}

func RegisterMovieService(e *echo.Echo) {
	handler := NewMovieHandler()
	e.GET("/movies", handler.MovieList)

}

func (m *MovieHandler) MovieList(c echo.Context) error {
	movies := m.svc.GetAllMovies()
	return c.JSON(http.StatusOK, movies)
}
