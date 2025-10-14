package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Assignment_id    uint           `json:"assigment_id"`
	Assignment       Assignment     `gorm:"foreignKey:Assignment_id"`
	ScheduleResponse datatypes.JSON `json:"schedule_response" gorm:"type:jsonb"`
}

type ScheduleResponse struct {
	ID             uint        `json:"id" gorm:"primaryKey"`
	Bestgeneration BestGenomes `json:"bestGeneration" gorm:"type:jsonb"`
	BestFitness    int         `json:"bestFitness"`
	Iteration      int         `json:"iteration"`
	Time           int         `json:"time"`
	ScheduleId     uint        `json:"scheduleId"`
}

type BestGenome struct {
	ID        uint   `json:"id"`
	DayIndex  uint   `json:"dayIndex"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Subject   string `json:"subject"`
	Room      string `json:"room"`
	Teachers  string `json:"teacher"`
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
