package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Template datatypes.JSON `json:"template"`

	Subjects datatypes.JSON `json:"subjects"`
	Teachers datatypes.JSON `json:"teachers"`
	Rooms    datatypes.JSON `json:"rooms"`
	Groups   datatypes.JSON `json:"groups"`

	ProcessName string  `json:"processName"`
	Population  int     `json:"population"`
	Generations int     `json:"generations"`
	Mutation    float64 `json:"mutation"`
	CrossOver   float64 `json:"cross_over"`
	Selection   float64 `json:"selction"`
	Reinsertion float64 `json:"reinsertion"`
}

type Process struct {
	SelectedData SelectedData `json:"selectedData"`
	ProcessData  ProcessData  `json:"processData"`
}

type SelectedData struct {
	SelectedTemplate uint         `json:"selectedTemplate"`
	SelectedSubjects []Subject    `json:"selectedSubjects" gorm:"many2many:assignment_subjects"`
	SelectedRooms    []Room       `json:"selectedRooms" gorm:"many2many:assignment_teachers"`
	SelectedTeachers []Teacher    `json:"selectedTeachers" gorm:"many2many:assignment_rooms"`
	SelectedGroups   []GroupInput `json:"selectedGroups" gorm:"many2many:assignment_groups"`
}

type ProcessData struct {
	ProcessName string  `json:"processName"`
	Population  int     `json:"population"`
	Generations int     `json:"generations"`
	Mutation    float64 `json:"mutation"`
	CrossOver   float64 `json:"cross_over"`
	Selection   float64 `json:"selection"`
	Reinsertion float64 `json:"reinsertion"`
}
