package services

import (
	"github.com/SeriesCity/Gateway/infrastructure/repository"
	"github.com/SeriesCity/Gateway/internal/core/entities"
	"github.com/SeriesCity/Gateway/internal/core/ports"
)

type SerialService struct {
	database ports.SerialServiceContract
}

func NewSerialService() *SerialService {
	db := repository.NewGormDatabase()
	return &SerialService{
		database: db,
	}
}

func (m *SerialService) GetAllSerials() []entities.Serial {
	return m.database.GetAllSerials()
}

func (m *SerialService) GetSingleSerial(slug string) (entities.Serial, error) {
	serial, err := m.database.GetSingleSerial(slug)
	return serial, err
}
