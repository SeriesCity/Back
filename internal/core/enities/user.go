package enities

import "time"

type UserPermission int16

const (
	RegularUser UserPermission = iota
	VIPUser
	SuperUser
	Admin
)

type User struct {
	FirstName      string
	LastName       string
	PhoneNumber    string
	Username       string
	Email          string `gorm:""`
	Password       string
	Collections    []Collection
	Permission     UserPermission `gorm:""`
	PermissionDate time.Time
}

type Collection struct {
	Name      string
	OwnerID   User
	Movies    []Movie
	Serials   []Serial
	IsPrivate bool `gorm:"default:false"`
}
