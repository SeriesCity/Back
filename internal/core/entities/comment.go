package entities

type Comment struct {
	BaseModel
	OwnerID    uint
	Owner      User `gorm:"foreignKey:OwnerID"`
	ParentID   uint
	Children   []Comment `gorm:"foreignKey:ParentID"`
	Content    string
	IsAccepted bool
	IsSpoiler  bool
	Likes      []User `gorm:"many2many:comment_likes;" json:"likes"`
	Dislikes   []User `gorm:"many2many:comment_dislikes;" json:"dislikes"`
}
