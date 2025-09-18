package model

import "gorm.io/gorm"

type GroupInput struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	SubjectIDs []uint `json:"subjects"` // Cambia a array de IDs
}

type Group struct {
	gorm.Model
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	Subjects []Subject `gorm:"many2many:group_subjects;"`
}

type group_subject struct {
	subject_id uint
	group_id   uint
}
