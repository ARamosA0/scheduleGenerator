package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Assignment_id uint       `json:"assigment_id"`
	Assignment    Assignment `gorm:"foreignKey:AssignmentID"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       time.Time  `json:"end_date"`
	Title         string     `json:"title"`
	Tooltip       string     `json:"tooltip"`
}
