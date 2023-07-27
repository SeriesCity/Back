package entities

type Serial struct {
	FilmBase
	CurrentStatus    string    `gorm:""`
	CreatorCountries []string  `gorm:""`
	Sessions         []Session `gorm:""`
}
