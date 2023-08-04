package ports

import "github.com/SeriesCity/Gateway/internal/core/entities"

type SerialServiceContract interface {
	GetAllSerials() []entities.Serial
	GetSingleSerial(slug string) (entities.Serial, error)
}
type SerialHandlerContract interface {
	GetAllSerials() []entities.Serial
	GetSingleSerial(slug string) (entities.Serial, error)
}
