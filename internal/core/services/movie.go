package services

import (
	"github.com/SeriesCity/Gateway/infrastructure/repository"
	"github.com/SeriesCity/Gateway/internal/core/entities"
	"github.com/SeriesCity/Gateway/internal/core/ports"
)

type MovieService struct {
	database ports.MovieServiceContract
}

func NewMovieService() *MovieService {
	db := repository.NewGormDatabase()
	return &MovieService{
		database: db,
	}
}

func (m *MovieService) GetAllMovies() []entities.Movie {
	return m.database.GetAllMovies()
}
