package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Assignment_id     uint        `json:"assigment_id"`
	Assignment        Assignment  `gorm:"foreignKey:Assignment_id"`
	ScheduleResponses BestGenomes `json:"schedule_response" gorm:"type:json"`
}

type ScheduleResponse struct {
	ID             uint         `json:"id"`
	Bestgeneration []BestGenome `json:"bestGeneration"`
	BestFitness    int          `json:"bestFitness"`
	Iteration      int          `json:"iteration"`
	ScheduleId     uint         `json:"scheduleId"`
}

type BestGenome struct {
	ID        uint   `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Title     string `json:"title"`
	Tooltip   string `json:"tooltip"`
}

type BestGenomes []BestGenome

func (s BestGenomes) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *BestGenomes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source was not []byte")
	}
	return json.Unmarshal(bytes, s)
}
