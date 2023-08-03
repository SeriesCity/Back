package main

import (
	"github.com/SeriesCity/Gateway/controller"
	"github.com/SeriesCity/Gateway/infrastructure/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	_ = godotenv.Load()
	e := echo.New()
	_, err := repository.GormInit()
	if err != nil {
		log.Fatal(err)
	}

	controller.RegisterMovieService(e)
	log.Fatal(e.Start(":8000"))
}
