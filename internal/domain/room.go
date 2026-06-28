package domain

type Room struct {
	BaseModel
	Name string `json:"name"`
	Type string `json:"type"`
}

type RoomMember struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}
