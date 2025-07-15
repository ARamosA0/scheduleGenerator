package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	Name      string         `json:"name"`
	DaysRange datatypes.JSON `json:"name"`
}

type TimeRange struct {
	Day       string `json:"day"`
	StartHour string `json:"startHour"`
	EndHour   string `json:"endHour"`
	Period    string `json:"period"`
	Status    bool   `json:"status"`
}
