package request

import (
	"github.com/google/uuid"
	"go-gorm/entity"
)

type RoomRequest struct {
	ID          string `json:"id" validate:"omitempty,uuid"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty,alphanum"`
	MaxMembers  int    `json:"maxMembers" validate:"omitempty"`
}

func (r *RoomRequest) ToDto() *entity.RoomDTO {
	var id uuid.UUID
	if r.ID != "" {
		id = uuid.Must(uuid.Parse(r.ID))
	}

	return &entity.RoomDTO{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		MaxMembers:  r.MaxMembers,
	}
}

type RoomQueryRequest struct {
	ID string `query:"id" param:"id" validate:"required,uuid"`
}
