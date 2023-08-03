package entities

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model
}

type PhotoFormat string

const (
	JPEG PhotoFormat = "JPEG"
	PNG  PhotoFormat = "PNG"
	GIF  PhotoFormat = "GIF"
)

type VideoFormat string

const (
	MP4 VideoFormat = "MP4"
	AVI VideoFormat = "AVI"
	MKV VideoFormat = "MKV"
)

type FilmBase struct {
	BaseModel
	Slug          string `gorm:"uniqueIndex"`
	Title         string
	Category      []Category `gorm:"many2many:movie_categories;" json:"genres"`
	PublishYear   time.Time
	Imdb          IMDB      `gorm:"embedded"`
	Comments      []Comment `gorm:"many2many:comments;" json:"comments"`
	Quality       string
	Duration      int32
	PermissionAge string
	Streamer      string
	Directors     []string `gorm:"type:text[]"`
	Writers       []string `gorm:"type:text[]"`
	Actors        []Actor  `gorm:"many2many:film_actors;" json:"actors"`
	Description   string
	Rates         []Rate `gorm:"many2many:film_rates;" json:"rates"`
	Scores        int16
	Thumbnail     Photo    `gorm:"embedded"`
	Cover         Photo    `gorm:"embedded"`
	Trailer       Video    `gorm:"embedded"`
	Tags          []string `gorm:"type:text[]"`
}

type IMDB struct {
	Link  string
	Rate  float32
	Votes int64
}

type ActorPerson struct {
	BaseModel
	FirstName   string
	LastName    string
	DateOfBirth string
	Photo       Photo `gorm:"embedded"`
	Profession  string
}

type Actor struct {
	BaseModel
	PersonID    uint
	Person      ActorPerson `gorm:"foreignKey:PersonID"`
	NameInMovie string
}

type Rate struct {
	BaseModel
	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`
	Score   int8 `gorm:"default:0"`
}

type DownloadLink struct {
	BaseModel
	Episode       uint  `gorm:"default:0"`
	Video         Video `gorm:"embedded"`
	Volume        string
	SampleQuality Photo `gorm:"embedded"`
	Quality       string
	Encode        string
	IsSoftSub     string `gorm:"default:false;"`
	IsDouble      bool   `gorm:"default:false;"`
	VIP           bool   `gorm:"default:false;"`
	OnlineStream  bool   `gorm:"default:false;"`
}
