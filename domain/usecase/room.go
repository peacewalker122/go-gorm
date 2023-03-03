package usecase

import (
	"context"
	"github.com/google/uuid"
	"go-gorm/entity"
)

type RoomUsecase interface {
	CreateRoom(ctx context.Context, room *entity.RoomDTO) error
	GetRoomByID(ctx context.Context, roomId, userID uuid.UUID) (*entity.RoomDTO, error)
	UpdateRoom(ctx context.Context, userID uuid.UUID, room *entity.RoomDTO) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}
