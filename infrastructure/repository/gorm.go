package repository

import (
	"fmt"
	"github.com/SeriesCity/Gateway/internal/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Postgres struct {
	db *gorm.DB
}

func GormInit() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DB_NAME")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", host, user, password, dbName, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(
		&entities.Category{},
		&entities.User{},
		&entities.Comment{},
		&entities.DownloadLink{},
		&entities.Session{},
		&entities.Rate{},
		&entities.ActorPerson{},
		&entities.Actor{},
		&entities.Movie{},
		&entities.Serial{},
		&entities.Collection{},
	)
	if err != nil {
		fmt.Println(err)
	}
	return database, nil
}
