package repository

import (
	"context"
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
)

type Room interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoomByID(ctx context.Context, roomId, userID uuid.UUID, opt ...string) (*entity.RoomDTO, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}
