package services

import (
	"time"

	"backend.com/backend/internal/model"
)

var teacher = []model.Teacher{
	{
		ID:             1,
		Name:           "Teacher1",
		AvailableTimes: []time.Time{time.Now()}, // Use slice initialization
	},
	{
		ID:             2,
		Name:           "Teacher2",
		AvailableTimes: []time.Time{time.Now()}, // Use slice initialization
	},
}

func GetAllTeachers() []model.Teacher {
	return teacher
}
