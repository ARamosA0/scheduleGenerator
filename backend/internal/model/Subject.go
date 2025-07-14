package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Code               string   `json:"code"`
	Name               string   `json:"name"`
	Credits            int      `json:"credits"`
	Hours              int      `json:"hours"`
	Semester           int      `json:"semester"`
	Career             string   `json:"career"`
	Requirements       string   `json:"requirements"`
	Description        string   `json:"description"`
	RequiredRoomTypeID int      `json:"required_room_type_id"`
	RequiredRoomType   RoomType `gorm:"foreignKey:RequiredRoomTypeID"`
}
