package entity

import (
	"errors"
	"github.com/google/uuid"
	"go-gorm/pkg/enum"
	"go-gorm/pkg/util"
	"time"
)

type TicketDTO struct {
	ID        uuid.UUID         `json:"id"`
	Metadata  string            `json:"metadata"`
	RoomID    uuid.UUID         `json:"roomId"`
	IssuedBy  uuid.UUID         `json:"issuedBy"`
	Assigned  uuid.UUID         `json:"assigned"`
	Status    enum.TicketStatus `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func (t *TicketDTO) Update(new *TicketDTO) {
	if new.Assigned != uuid.Nil {
		t.Assigned = new.Assigned
	}
	if new.Metadata != "" {
		t.Metadata = new.Metadata
	}
	if new.Status != "" {
		t.Status = new.Status
	}
}

func (t *TicketDTO) Validate() error {
	var multierror util.MultiError

	if t.IssuedBy == uuid.Nil {
		multierror.Add(errors.New("issuedBy is empty"))
	}
	if t.Metadata == "" {
		multierror.Add(errors.New("metadata is empty"))
	}

	if multierror.HasError() {
		return &multierror
	}

	return nil
}

type TicketFilter struct {
	RoomID     uuid.UUID         `json:"roomId"`
	Status     enum.TicketStatus `json:"status"`
	IssuedBy   uuid.UUID         `json:"issuedBy"`
	Assigned   uuid.UUID         `json:"assigned"`
	TimeFilter time.Time         `json:"timeFilter"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
}

func (t *TicketFilter) Validate() error {
	var multierror util.MultiError

	if t.RoomID == uuid.Nil {
		multierror.Add(errors.New("roomId is empty"))
	}
	if t.IssuedBy == uuid.Nil {
		multierror.Add(errors.New("issuedBy is empty"))
	}
	if t.Status == "" {
		t.Status = enum.Todo
	}
	if multierror.HasError() {
		return &multierror
	}

	return nil
}
