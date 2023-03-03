package model

import "github.com/google/uuid"

type UserRoom struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	RoomID uuid.UUID `gorm:"primaryKey"`
}

func (UserRoom) TableName() string {
	return "user_room"
}
