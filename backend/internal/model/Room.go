package model

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	RoomType int    `json:"room_type"`
}

type RoomType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
