package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type Category struct {
	CategoryId int    `gorm:"type:tinyint;primaryKey;index:,unique"`
	Name       string `gorm:"type:varchar(25)"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRequest struct {
	*Category
}

type CategoryResponse struct {
	*Category
}

func NewCategoryResponse(category *Category) *CategoryResponse {
	return &CategoryResponse{category}
}

func NewCategoryListResponse(cs []*Category) []render.Renderer {
	list := []render.Renderer{}
	for _, c := range cs {
		list = append(list, NewCategoryResponse(c))
	}
	return list
}

func (a *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
