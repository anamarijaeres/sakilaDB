package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type Address struct {
	Address   string `gorm:"type:varchar(50)"`
	AddressId int    `gorm:"type:smallint;primaryKey"`
	Phone     string `gorm:"type:varchar(20)"`
}

func (Address) TableName() string {
	return "address"
}

type AddressRequest struct {
	*Address
}

type AddressResponse struct {
	*Address
}

func NewAddressResponse(address *Address) *AddressResponse {
	return &AddressResponse{address}
}

func NewAddressListResponse(addresses []*Address) []render.Renderer {
	list := []render.Renderer{}
	for _, ad := range addresses {
		list = append(list, NewAddressResponse(ad))
	}
	return list
}

func (a *AddressResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
