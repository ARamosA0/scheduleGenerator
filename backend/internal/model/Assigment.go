package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// type Assignment struct {
// 	gorm.Model
// 	SubjectID    uint    `json:"subject_id"`
// 	Subject      Subject `gorm:"foreignKey:SubjectID"`
// 	TeacherID    uint    `json:"teacher_id"`
// 	Teacher      Teacher `gorm:"foreignKey:TeacherID"`
// 	GroupID      uint    `json:"group_id"`
// 	Group        Group   `gorm:"foreignKey:GroupID"`
// 	RoomID       uint    `json:"room_id"`
// 	Room         Room    `gorm:"foreignKey:RoomID"`
// 	HoursPerWeek int     `json:"hours_per_week"`
// }

type Assignment struct {
	gorm.Model
	TemplateID uint     `json:"templateId"`
	Template   Template `gorm:"foreignKey:TemplateID"`

	Subjects datatypes.JSON `json:"subjects"`
	Teachers datatypes.JSON `json:"teachers"`
	Rooms    datatypes.JSON `json:"rooms"`

	ProcessName string  `json:"processName"`
	Population  int     `json:"population"`
	Generations int     `json:"generations"`
	Mutation    float64 `json:"mutation"`
	CrossOver   float64 `json:"crossOver"`
	Elitism     float64 `json:"elitism"`
}

type Process struct {
	SelectedData SelectedData `json:"selectedData"`
	ProcessData  ProcessData  `json:"processData"`
}

type SelectedData struct {
	SelectedTemplate uint      `json:"selectedTemplate"`
	SelectedSubjects []Subject `json:"selectedSubjects" gorm:"many2many:assignment_subjects"`
	SelectedRooms    []Room    `json:"selectedRooms" gorm:"many2many:assignment_teachers"`
	SelectedTeachers []Teacher `json:"selectedTeachers" gorm:"many2many:assignment_rooms"`
}

type ProcessData struct {
	ProcessName string  `json:"processName"`
	Population  int     `json:"population"`
	Generations int     `json:"generation"`
	Mutation    float64 `json:"mutation"`
	CrossOver   float64 `json:"crossOver"`
	Elitism     float64 `json:"elitism"`
}
