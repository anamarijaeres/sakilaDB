package models

type Category struct {
	CategoryId int `gorm:"type:tinyint;primaryKey;index:,unique"`
	Name       int `gorm:"type:varchar(25)"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRequest struct {
	*Category
}
