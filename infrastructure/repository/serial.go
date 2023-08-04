package repository

import (
	"errors"
	"github.com/SeriesCity/Gateway/internal/core/entities"
)

func (p *PGRepository) GetAllSerials() []entities.Serial {
	var serials []entities.Serial
	p.DB.Model(entities.Serial{}).Find(&serials)
	return serials
}

func (p *PGRepository) GetSingleSerial(slug string) (entities.Serial, error) {
	var serial entities.Serial
	found := p.DB.Model(entities.Serial{}).Where("slug = ?", slug).First(&serial).RowsAffected
	if found != 1 {
		return serial, errors.New("slug not found")
	}
	return serial, nil
}
