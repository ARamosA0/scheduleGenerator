package model

type Assignment struct {
	ID           int `json:"id"`
	SubjectID    int `json:"subject_id"`
	TeacherID    int `json:"teacher_id"`
	GroupID      int `json:"group_id"`
	RoomID       int `json:"room_id"`
	HoursPerWeek int `json:"hours_per_week"`
}
