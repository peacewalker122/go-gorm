package repository

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
)

type User interface {
	GetByID(id uuid.UUID) (*entity.UserDTO, error)
	GetUser(val interface{}, opt ...string) (*entity.UserDTO, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
