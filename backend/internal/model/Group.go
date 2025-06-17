package model

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Semester string `json:"semester"`
	Size     int    `json:"size"`
}
