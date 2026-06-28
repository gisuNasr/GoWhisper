package domain

type Message struct {
	BaseModel
	RoomID           int64  `json:"room_id"`
	UserID           int64  `json:"user_id"`
	DeviceID         int64  `json:"device_id"`
	EncryptedPayload string `json:"encrypted_payload"`
}
