package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
)

func (u *Usecase) CreateRoom(ctx context.Context, room *entity.RoomDTO) error {
	if room == nil {
		return errors.New("nil room")
	}

	if err := room.Validate(); err != nil {
		return err
	}
	id := u.uuidfn()
	roomModel := model.RoomToModel(id, room)
	return u.Room.CreateRoom(ctx, roomModel)
}

func (u *Usecase) GetRoomByID(ctx context.Context, roomID, userID uuid.UUID) (*entity.RoomDTO, error) {
	if roomID == uuid.Nil || userID == uuid.Nil {
		return nil, errors.New("id is empty")
	}

	return u.Room.GetRoomByID(ctx, roomID, userID)
}

func (u *Usecase) UpdateRoom(ctx context.Context, userID uuid.UUID, newRoom *entity.RoomDTO) error {
	if newRoom == nil {
		return errors.New("room is empty")
	}

	oldRoom, err := u.Room.GetRoomByID(ctx, newRoom.ID, userID, "update")
	if err != nil {
		return err
	}
	oldRoom.Update(newRoom)

	if err := oldRoom.Validate(); err != nil {
		return err
	}

	return u.Room.UpdateRoom(ctx, model.RoomToModel(u.uuidfn(), oldRoom))
}

func (u *Usecase) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("id is empty")
	}

	return u.Room.DeleteRoom(ctx, id)
}
