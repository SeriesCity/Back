package repository

import "github.com/SeriesCity/Gateway/internal/core/entities"

func (p *PGRepository) GetAllMovies() []entities.Movie {
	var movies []entities.Movie
	p.DB.Model(entities.Movie{}).Find(&movies).Order("-updatedAt")
	return movies
}
