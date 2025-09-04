package model

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	Name      string         `json:"name"`
	DaysRange datatypes.JSON `json:"daysRange"`
}

type TimeRange struct {
	Day       int      `json:"day"`
	StartHour TimeOnly `json:"startHour"`
	EndHour   TimeOnly `json:"endHour"`
	Period    int      `json:"period"`
	Status    bool     `json:"status"`
}

type TimeOnly time.Time

func (t *TimeOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsed, err := time.Parse("15:04", s)
	if err != nil {
		return err
	}
	*t = TimeOnly(parsed)
	return nil
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04"))), nil
}
