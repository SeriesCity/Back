package enities

type Movie struct {
	FilmBase
	Links []DownloadLink `gorm:""`
}
