package models

import "time"

type Inventory struct {
	FilmId      int       `gorm:"type:smallint;primaryKey"`
	InventoryId int       `gorm:"type:mediumint"`
	LastUpdate  time.Time `gorm:"type:autoCreateTime"`
	StoreId     int       `gorm:"type:tinyint;primaryKey"`
}

func (Inventory) TableName() string {
	return "inventory"
}

type InventoryRequest struct {
	*Inventory
}
