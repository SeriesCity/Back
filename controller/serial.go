package controller

import (
	"github.com/SeriesCity/Gateway/internal/core/ports"
	"github.com/SeriesCity/Gateway/internal/core/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SerialHandler struct {
	svc ports.SerialHandlerContract
}

func NewSerialHandler() *SerialHandler {
	svc := services.NewSerialService()
	return &SerialHandler{
		svc: svc,
	}
}

func RegisterSerialService(e *echo.Echo) {
	handler := NewSerialHandler()
	group := e.Group("/serial")
	group.GET("/", handler.SerialList)
	group.GET("/:slug/", handler.GetSingleSerial)

}

func (m *SerialHandler) SerialList(c echo.Context) error {
	movies := m.svc.GetAllSerials()
	return c.JSON(http.StatusOK, movies)
}

func (m *SerialHandler) GetSingleSerial(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg": "slug required",
		})
	}
	movies, err := m.svc.GetSingleSerial(slug)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, movies)
}
