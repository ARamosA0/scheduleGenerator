package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name          string        `json:"name"`
	LastName      string        `json:"lastName"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	Specialty     int           `json:"specialty"`
	AvailableDays AvailableDays `gorm:"type:jsonb" json:"available_days"`
	// AvailableTimes TimesWrapper `gorm:"type:jsonb" json:"available_times"`
}

type AvailableDays []string

func (ad AvailableDays) Value() (driver.Value, error) {
	return json.Marshal(ad)
}

func (ad *AvailableDays) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Expected []byte, got %T", value)
	}
	return json.Unmarshal(bytes, ad)
}

// type TimesWrapper []time.Time

// func (tw TimesWrapper) Value() (driver.Value, error) {
// 	return json.Marshal(tw)
// }

// func (tw *TimesWrapper) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return fmt.Errorf("Expected []byte, got %T", value)
// 	}
// 	return json.Unmarshal(bytes, tw)
// }
