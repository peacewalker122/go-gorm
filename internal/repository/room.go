package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-gorm/domain/repository"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
	"go-gorm/pkg/util"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
	tx *util.DatabaseTx
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *model.Room) error {
	var err error
	r.tx.Tx(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).
			Create(room).
			Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).
			Create(room.ToUserRoomModel(room.CreatedBy)).
			Error
		if err != nil {
			return err
		}
		return err
	})

	return err
}

func (r *RoomRepository) JoinRoom(ctx context.Context, userID, roomID uuid.UUID) error {
	var err error

	err = r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Where("room_id = ?", roomID).
		First(&model.UserRoom{}).Error

	if err == nil {
		return errors.New("user already in the room")
	}

	r.tx.Tx(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Create(&model.UserRoom{
			UserID: userID,
			RoomID: roomID,
		}).Error
		if err != nil {
			return err
		}
		return err
	})

	return err
}

func (r *RoomRepository) GetRoomByID(ctx context.Context, roomId, userID uuid.UUID, opt ...string) (*entity.RoomDTO, error) {
	var (
		room  = &model.Room{}
		err   error
		query = r.db
	)

	if len(opt) > 0 {
		switch opt[0] {
		case "update":
			query = query.Model(room).
				WithContext(ctx).
				Select("*").
				Where(
					r.db.Where("id = ?", roomId).
						Where("created_by = ?", userID),
				).
				Scan(room)
		}
	} else {
		query = query.
			WithContext(ctx).
			Model(room).
			Joins("JOIN user_room ON user_room.room_id = room.id").
			Where(
				r.db.Where("room.id = ?", roomId.String()).
					Where("user_room.user_id = ?", userID.String()),
			).
			Scan(room)
	}

	err = query.Error
	if err != nil {
		return nil, err
	}

	return room.ToDTO(), nil
}

func (r *RoomRepository) UpdateRoom(ctx context.Context, room *model.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	err := r.db.Update("deleted_at", gorm.Expr("now()")).Where("id = ?", id).Error

	return err
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
		tx: util.NewDatabaseTx(db),
	}
}

var _ repository.Room = (*RoomRepository)(nil)
