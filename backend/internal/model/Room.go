package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	RoomType int    `json:"room_type"`
}

type RoomType struct {
	gorm.Model
	Name string `json:"name"`
}
