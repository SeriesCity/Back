package ports

import "github.com/SeriesCity/Gateway/internal/core/entities"

type CategoryServiceContract interface {
	GetCategories() []entities.Category
	CreateNewCategory(categoryName string) (entities.Category, error)
}

type CategoryHandlerContract interface {
	GetAllCategories() []entities.Category
	CreateNewCategory(categoryName string) (entities.Category, error)
}
