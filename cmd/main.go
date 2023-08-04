package main

import (
	"github.com/SeriesCity/Gateway/controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	_ = godotenv.Load()
	e := echo.New()
	controller.RegisterMovieService(e)
	controller.RegisterSerialService(e)
	log.Fatal(e.Start(":8000"))
}
