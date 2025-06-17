package model

type Subject struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Credits          string `json:"credits"`
	RequiredRoomType int    `json:"required_room_type"`
}
