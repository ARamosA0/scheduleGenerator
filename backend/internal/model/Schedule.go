package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Assignment_id     uint              `json:"assigment_id"`
	Assignment        Assignment        `gorm:"foreignKey:Assignment_id"`
	ScheduleResponses ScheduleResponses `json:"schedule_response" gorm:"type:json"`
	// StartDate     time.Time  `json:"start_date"`
	// EndDate       time.Time  `json:"end_date"`
	// Title         string     `json:"title"`
	// Tooltip       string     `json:"tooltip"`
}

type ScheduleResponse struct {
	ID        uint   `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Title     string `json:"title"`
	Tooltip   string `json:"tooltip"`
}

type ScheduleResponses []ScheduleResponse

func (s ScheduleResponses) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *ScheduleResponses) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source was not []byte")
	}
	return json.Unmarshal(bytes, &s)
}
