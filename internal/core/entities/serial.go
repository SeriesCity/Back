package entities

type Serial struct {
	FilmBase
	CurrentStatus    string
	CreatorCountries []string  `gorm:"type:text[]"`
	Sessions         []Session `gorm:"many2many:sessions;" json:"sessions"`
}

type Session struct {
	BaseModel
	Number   int8
	Episodes []DownloadLink `gorm:"many2many:episodes_links;" json:"episodes"`
}
