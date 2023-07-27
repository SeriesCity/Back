package enities

type Serial struct {
	FilmBase
	CurrentStatus    string    `gorm:""`
	CreatorCountries []string  `gorm:""`
	Sessions         []Session `gorm:""`
}
