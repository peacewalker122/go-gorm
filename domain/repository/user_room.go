package repository

import "github.com/google/uuid"

type UserRoom interface {
	IsUserInRoom(userID, roomID uuid.UUID) (bool, error)
}
