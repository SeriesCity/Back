package controller

import (
	"github.com/SeriesCity/Gateway/internal/core/ports"
	"github.com/SeriesCity/Gateway/internal/core/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryHandler struct {
	svc ports.CategoryHandlerContract
}

func NewCategoryHandler() *CategoryHandler {
	svc := services.NewCategoryService()
	return &CategoryHandler{
		svc: svc,
	}
}

func RegisterCategoryService(e *echo.Echo) {
	handler := NewCategoryHandler()
	group := e.Group("/category")
	group.GET("/", handler.Categories)
	group.POST("/", handler.CreateCategory)

}

func (h *CategoryHandler) Categories(c echo.Context) error {
	categories := h.svc.GetAllCategories()
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	categoryName := c.FormValue("name")
	category, err := h.svc.CreateNewCategory(categoryName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, category)
}
