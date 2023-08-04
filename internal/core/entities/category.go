package entities

type Category struct {
	BaseModel
	Name string `gorm:"Column:category_name; unique"`
}
