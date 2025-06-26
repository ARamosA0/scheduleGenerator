package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name       string   `json:"name"`
	Capacity   string   `json:"capacity"`
	RoomTypeID uint     `json:"room_type_id"`
	RoomType   RoomType `gorm:"foreignKey:RoomTypeID"`
}

type RoomType struct {
	gorm.Model
	Name string `json:"name"`
}
