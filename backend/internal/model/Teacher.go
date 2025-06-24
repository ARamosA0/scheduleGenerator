package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name           string       `json:"name"`
	AvailableTimes TimesWrapper `gorm:"type:jsonb" json:"available_times"`
}

type TimesWrapper []time.Time

func (tw TimesWrapper) Value() (driver.Value, error) {
	return json.Marshal(tw)
}

func (tw *TimesWrapper) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Expected []byte, got %T", value)
	}
	return json.Unmarshal(bytes, tw)
}
