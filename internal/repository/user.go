package repository

import (
	"github.com/google/uuid"
	"go-gorm/domain/repository"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
	"gorm.io/gorm"
	"reflect"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetByID(id uuid.UUID) (*entity.UserDTO, error) {
	var user model.User
	err := u.db.First(&user, id).Where("deleted_at", nil).Error
	if err != nil {
		return nil, err
	}
	return user.ToDTO(), nil
}

func (u *UserRepository) GetUser(val interface{}, opt ...string) (*entity.UserDTO, error) {
	var (
		user     model.User
		subQuery = u.db.Model(&user).Where("deleted_at", nil)
	)

	if types := reflect.TypeOf(val); types.Kind() == reflect.String {
		if opt[0] == "username" {
			subQuery.Where("username = ?", val)
		}
		if opt[0] == "email" {
			subQuery.Where("email = ?", val)
		}
	}

	err := u.db.Where(subQuery).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user.ToDTO(), nil
}

func (u *UserRepository) Create(user *model.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Update(user *model.User) error {
	err := u.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	err := u.db.Update("deleted_at", gorm.Expr("NOW()")).Where("id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) repository.User {
	return &UserRepository{db: db}
}

// Path: internal/repository/user.go
var _ repository.User = &UserRepository{}
