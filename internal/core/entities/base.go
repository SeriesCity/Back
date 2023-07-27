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
	Slug          string     `gorm:"uniqueIndex"`
	Title         string     `gorm:""`
	Category      []string   `gorm:""`
	PublishYear   *time.Time `gorm:""`
	IMDB          IMDB       `gorm:""`
	Comments      []Comment  `gorm:""`
	Quality       string     `gorm:""`
	Duration      int32      `gorm:""`
	PermissionAge string     `gorm:""` // enum
	Genres        []string   `gorm:""`
	Streamer      string     `gorm:""`
	Directors     []string   `gorm:""`
	Writers       []string   `gorm:""`
	Actors        []Actor    `gorm:""`
	Description   string     `gorm:""`
	Rates         []Rate     `gorm:""`
	Scores        int16      `gorm:""`
	Thumbnail     Photo      `gorm:""`
	Cover         Photo      `gorm:""`
	Trailer       Video      `gorm:""`
	Tags          []string   `gorm:""`
}

type Session struct {
	Number   int8           `gorm:""`
	Episodes []DownloadLink `gorm:""`
}

type IMDB struct {
	Link  string  `gorm:""`
	Rate  float32 `gorm:""`
	Votes int64   `gorm:""`
}

type ActorPerson struct {
	FirstName   string `gorm:""`
	LastName    string `gorm:""`
	DateOfBirth string `gorm:""`
	Photo       Photo  `gorm:""`
	Profession  string `gorm:""`
}

type Actor struct {
	Person      ActorPerson `gorm:""`
	NameInMovie string      `gorm:""`
}

type Rate struct {
	Owner User `gorm:""`
	Score int8 `gorm:""`
}

type DownloadLink struct {
	Format        string `gorm:""`
	Volume        string `gorm:""`
	SampleQuality Photo  `gorm:"embedded;embeddedPrefix:sample_"`
	Link          string `gorm:""`
	Quality       string `gorm:""`
	Encode        string `gorm:""`
	IsSoftSub     string `gorm:"default:false;"`
	IsDouble      bool   `gorm:"default:false;"`
	VIP           bool   `gorm:"default:false;"`
	OnlineStream  bool   `gorm:"default:false;"`
}
