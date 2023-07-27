package entities

type Movie struct {
	FilmBase
	Links []DownloadLink `gorm:""`
}
