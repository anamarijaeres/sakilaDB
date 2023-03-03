package models

type FilmCategory struct {
	CategoryId int `gorm:"type:tinyint;primaryKey;index:,unique"`
	FilmId     int `gorm:"type:smallint;primaryKey;index:,unique"`
}

func (FilmCategory) TableName() string {
	return "film_category"
}

type FilmCategoryRequest struct {
	*FilmCategory
}
