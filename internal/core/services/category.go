package services

import (
	customError "github.com/SeriesCity/Gateway/custom_error"
	"github.com/SeriesCity/Gateway/infrastructure/repository"
	"github.com/SeriesCity/Gateway/internal/core/entities"
	"github.com/SeriesCity/Gateway/internal/core/ports"
)

type CategoryService struct {
	database ports.CategoryServiceContract
}

func NewCategoryService() *CategoryService {
	db := repository.NewGormDatabase()
	return &CategoryService{
		database: db,
	}
}

func (c *CategoryService) GetAllCategories() []entities.Category {
	return c.database.GetCategories()
}

func (c *CategoryService) CreateNewCategory(categoryName string) (entities.Category, error) {
	category, err := c.database.CreateNewCategory(categoryName)
	if err != nil {
		customizedErr := customError.FindDBError(err, "category")
		return category, customizedErr
	}
	return category, err
}
