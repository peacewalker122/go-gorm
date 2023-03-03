package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type RoomDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CreatedBy   uuid.UUID `json:"createdBy"`
	Description string    `json:"description"`
	MaxMembers  int       `json:"maxMembers"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (r *RoomDTO) Update(new *RoomDTO) {
	if new.Name != "" {
		r.Name = new.Name
	}
	r.Description = new.Description
	r.MaxMembers = new.MaxMembers
}

func (r *RoomDTO) Validate() error {
	if r.Name == "" {
		return errors.New("name is empty")
	}
	return nil
}
