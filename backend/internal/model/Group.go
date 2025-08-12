package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	Subjects []Subject `gorm:"many2many:group_subjects;"`
}
