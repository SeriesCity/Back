package entities

import (
	"time"
)

type UserPermission int16

const (
	RegularUser UserPermission = iota + 1
	VIPUser
	SuperUser
	Admin
)

type User struct {
	BaseModel
	FirstName      string    `gorm:"type:character" json:"first_name"`
	LastName       string    `gorm:"type:character" json:"last_name"`
	Username       string    `gorm:"type:character;not null" json:"username"`
	Email          string    `gorm:"type:character;uniqueIndex;not null" json:"email"`
	Phone          string    `gorm:"type:character" json:"phone"`
	Password       []string  `gorm:"type:text[]" json:"password"`
	Permission     int       `json:"permission"`
	PermissionDate time.Time `gorm:"Column:permission_date"`
}

type Collection struct {
	BaseModel
	Name      string
	Movies    []Movie  `gorm:"many2many:collection_movies;" json:"movies"`
	Serials   []Serial `gorm:"many2many:collection_serials;" json:"serials"`
	OwnerID   uint
	Owner     User `gorm:"foreignKey:OwnerID"`
	IsPrivate bool `gorm:"default:false"`
}
