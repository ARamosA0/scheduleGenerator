package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Credits          int     `json:"credits"`
	Hours            int     `json:"hours"`
	Semester         int     `json:"semester"`
	Career           string  `json:"career"`
	Requirements     string  `json:"requirements"`
	Description      string  `json:"description"`
	RequiredRoomType int     `json:"required_room_type"`
	Groups           []Group `gorm:"many2many:group_subjects;"`
}
