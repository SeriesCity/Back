package entities

import "time"

type Photo struct {
	Link   string      `gorm:""`
	Format PhotoFormat `gorm:""` // enum
}

type Video struct {
	Link     string         `gorm:""`
	Duration *time.Duration `gorm:""`
	Format   VideoFormat
}
