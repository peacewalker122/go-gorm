package model

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"time"
)

type Room struct {
	ID          uuid.UUID `gorm:"primaryKey default:uuid_generate_v4()"`
	Name        string    `gorm:"not null"`
	CreatedBy   uuid.UUID `gorm:"not null"`
	Description string    `gorm:"not null"`
	MaxMember   int       `gorm:"null"`
	CreatedAt   time.Time `gorm:"not null default:now()"`
}

func (Room) TableName() string {
	return "room"
}
func (r *Room) ToDTO() *entity.RoomDTO {
	return &entity.RoomDTO{
		ID:          r.ID,
		Name:        r.Name,
		CreatedBy:   r.CreatedBy,
		Description: r.Description,
		MaxMembers:  r.MaxMember,
		CreatedAt:   r.CreatedAt,
	}
}

func RoomToModel(id uuid.UUID, roomDto *entity.RoomDTO) *Room {

	if roomDto.ID == uuid.Nil {
		roomDto.ID = id
	}

	return &Room{
		ID:          roomDto.ID,
		Name:        roomDto.Name,
		CreatedBy:   roomDto.CreatedBy,
		Description: roomDto.Description,
		MaxMember:   roomDto.MaxMembers,
		CreatedAt:   roomDto.CreatedAt,
	}
}

func (r *Room) ToUserRoomModel(userid uuid.UUID) *UserRoom {
	return &UserRoom{
		UserID: userid,
		RoomID: r.ID,
	}
}
