package model

import "time"

type Teacher struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	AvailableTimes []time.Time `json:"available_times"`
}
