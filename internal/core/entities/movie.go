package entities

type Movie struct {
	FilmBase
	Links []DownloadLink `gorm:"many2many:movie_links;" json:"links"`
}
