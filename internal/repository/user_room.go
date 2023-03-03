package repository

import (
	"github.com/google/uuid"
	"go-gorm/internal/repository/model"
	"go-gorm/pkg/enum"
	"gorm.io/gorm"
)

type UserRoomRepository struct {
	db *gorm.DB
}

func NewUserRoomRepository(db *gorm.DB) *UserRoomRepository {
	return &UserRoomRepository{db: db}
}

func (u *UserRoomRepository) IsUserInRoom(userID, roomID uuid.UUID) (bool, error) {
	var (
		count      int64
		joinQuery  = u.db
		whereQuery = u.db
	)

	joinQuery = joinQuery.
		InnerJoins("inner join \"user\" u on u.id = user_room.user_id").
		Where(
			whereQuery.
				Where("u.id = ?", userID).
				Where("user_room.room_id = ?", roomID).
				Where(whereQuery.Where("u.deleted_at is null").Where("u.status = ?", enum.Active)),
		)

	err := joinQuery.
		Model(&model.UserRoom{}).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
