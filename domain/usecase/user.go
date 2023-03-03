package usecase

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/pkg/util"
)

type UserUsecase interface {
	CreateUser(user *entity.UserDTO) error
	GetUserByID(id uuid.UUID) (*entity.UserDTO, error)
	UpdateUser(user *entity.UserDTO) error
	Login(username, password string) (*entity.UserDTO, map[string]interface{}, *util.UsecaseError)
	DeleteUser(id uuid.UUID) error
}
