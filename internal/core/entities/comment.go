package entities

type Comment struct {
	Owner      User      `gorm:""` // user id
	Content    string    `gorm:""`
	IsAccepted bool      `gorm:""`
	IsSpoiler  bool      `gorm:""`
	Likes      []User    `gorm:""`
	Dislikes   []User    `gorm:""`
	Parent     []Comment `gorm:""`
}
