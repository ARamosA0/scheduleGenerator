package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name             string `json:"name"`
	Credits          string `json:"credits"`
	RequiredRoomType int    `json:"required_room_type"`
}
