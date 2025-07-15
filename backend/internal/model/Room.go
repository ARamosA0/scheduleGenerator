package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Code         string   `json:"code"`
	Name         string   `json:"name"`
	Capacity     string   `json:"capacity"`
	RoomTypeID   uint     `json:"room_type_id"`
	RoomType     RoomType `gorm:"foreignKey:RoomTypeID"`
	Floor        int      `json:"floor"`
	Building     string   `json:"building"`
	Observations string   `json:"observations"`
}

type RoomType struct {
	gorm.Model
	Name string `json:"name"`
}
