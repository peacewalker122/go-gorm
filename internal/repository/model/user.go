package model

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"time"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string     `json:"username" gorm:"unique;not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"password" gorm:"not null"`
	Status    string     `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp with time zone;default:now()"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:timestamp with time zone;default:null"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) ToDTO() *entity.UserDTO {
	return &entity.UserDTO{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
	}
}

func UserToModel(id uuid.UUID, userDto *entity.UserDTO) *User {

	if userDto.ID == uuid.Nil {
		userDto.ID = id
	}

	return &User{
		ID:        userDto.ID,
		Username:  userDto.Username,
		Email:     userDto.Email,
		Password:  userDto.Password,
		Status:    userDto.Status,
		CreatedAt: userDto.CreatedAt,
	}
}
