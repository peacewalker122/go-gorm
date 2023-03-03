package entity

import (
	"errors"
	"github.com/google/uuid"
	"go-gorm/pkg/util"
	"time"
)

const (
	ValidateCreate = "create"
	ValidateUpdate = "update"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserDTO) Update(user *UserDTO) {
	if user.Username != "" {
		u.Username = user.Username
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.Password != "" {
		u.Password = user.Password
	}
}

func (u *UserDTO) Validate(opt ...string) error {

	if u == nil {
		return errors.New("nil user")
	}
	if u.Username == "" {
		return errors.New("empty username")
	}
	if u.Email == "" {
		return errors.New("empty email")
	}
	if u.Password == "" {
		return errors.New("empty password")
	}

	if opt != nil && (opt[0] == "create" || opt[0] == "update") {
		tempPass, err := util.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = tempPass
	}

	return nil
}
