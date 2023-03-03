package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type Store struct {
	StoreId   int `gorm:"type:tinyint"`
	AddressId int `gorm:"type:smallint;primaryKey"`
}

func (Store) TableName() string {
	return "store"
}

type StoreRequest struct {
	*Store
}

type StoreResponse struct {
	*Store
}

func NewStoreResponse(store *Store) *StoreResponse {
	return &StoreResponse{store}
}

func NewStoreListResponse(stores []*Store) []render.Renderer {
	list := []render.Renderer{}
	for _, store := range stores {
		list = append(list, NewStoreResponse(store))
	}
	return list
}

func (a *StoreResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
