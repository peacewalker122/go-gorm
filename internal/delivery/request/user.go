package request

import (
	"github.com/google/uuid"
	"go-gorm/entity"
)

type LoginRequest struct {
	Username string `json:"username" validate:"omitempty,required,min=3,max=20"`
	Email    string `json:"email" validate:"omitempty,required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type UserRequest struct {
	ID       string `json:"id" query:"id" param:"id" validate:"omitempty,uuid"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

func (u *UserRequest) ToUserDto() *entity.UserDTO {
	var id uuid.UUID

	if u.ID != "" {
		id = uuid.Must(uuid.Parse(u.ID))
	}

	return &entity.UserDTO{
		ID:       id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

type QueryUserRequest struct {
	ID       string `json:"id" query:"id" param:"id" validate:"omitempty,uuid"`
	Username string `json:"username" query:"username" validate:"omitempty,min=3,max=20"`
	Email    string `json:"email" query:"username" validate:"omitempty,email"`
}
