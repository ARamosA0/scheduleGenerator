package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
