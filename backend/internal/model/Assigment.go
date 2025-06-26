package model

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	SubjectID    uint    `json:"subject_id"`
	Subject      Subject `gorm:"foreignKey:SubjectID"`
	TeacherID    uint    `json:"teacher_id"`
	Teacher      Teacher `gorm:"foreignKey:TeacherID"`
	GroupID      uint    `json:"group_id"`
	Group        Group   `gorm:"foreignKey:GroupID"`
	RoomID       uint    `json:"room_id"`
	Room         Room    `gorm:"foreignKey:RoomID"`
	HoursPerWeek int     `json:"hours_per_week"`
}
