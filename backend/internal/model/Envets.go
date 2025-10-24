package model

// "gorm.io/datatypes"
// "gorm.io/gorm"

type ProcessEventMessage struct {
	Message    string `json:"message"`
	Progress   uint   `json:"progress"`
	Status     string `json:"status"`
	ScheduleId string `json:"scheduleId"`
}
