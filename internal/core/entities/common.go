package entities

import "time"

type Photo struct {
	Link   string
	Format PhotoFormat
}

type Video struct {
	Link     string
	Duration *time.Duration
	Format   VideoFormat
}
