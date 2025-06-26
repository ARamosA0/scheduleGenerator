package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name               string   `json:"name"`
	Credits            string   `json:"credits"`
	RequiredRoomTypeID int      `json:"required_room_type_id"`
	RequiredRoomType   RoomType `gorm:"foreignKey:RequiredRoomTypeID"`
}
