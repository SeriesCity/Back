package repository

import (
	"github.com/SeriesCity/Gateway/internal/core/entities"
)

func (p *PGRepository) GetCategories() []entities.Category {
	var categories []entities.Category
	p.DB.Model(entities.Category{}).Find(&categories)
	return categories
}

func (p *PGRepository) CreateNewCategory(categoryName string) (entities.Category, error) {
	var category entities.Category
	category.Name = categoryName
	err := p.DB.Save(&category).Error
	return category, err
}
