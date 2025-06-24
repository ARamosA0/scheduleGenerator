package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name     string `json:"name"`
	Semester string `json:"semester"`
	Size     int    `json:"size"`
}
