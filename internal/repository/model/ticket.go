package model

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/pkg/enum"
	"time"
)

type Ticket struct {
	ID        uuid.UUID `gorm:"primaryKey default:uuid_generate_v4()"`
	Metadata  string
	RoomID    uuid.UUID         `gorm:"not null"`
	IssuedBy  uuid.UUID         `gorm:"not null"`
	Assigned  uuid.UUID         `gorm:"not null"`
	Status    enum.TicketStatus `gorm:"not null default:'todo'"`
	CreatedAt time.Time         `gorm:"not null default:now()"`
	UpdatedAt time.Time         `gorm:"null"`
}

func (t *Ticket) TableName() string {
	return "ticket"
}

func (t *Ticket) ToDTO() *entity.TicketDTO {
	return &entity.TicketDTO{
		ID:        t.ID,
		RoomID:    t.RoomID,
		Metadata:  t.Metadata,
		IssuedBy:  t.IssuedBy,
		Assigned:  t.Assigned,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func TicketToModel(id uuid.UUID, ticketDto *entity.TicketDTO) *Ticket {

	if ticketDto.ID == uuid.Nil {
		ticketDto.ID = id
	}
	if ticketDto.Status == "" {
		ticketDto.Status = enum.Todo
	}

	if ticketDto.Assigned == uuid.Nil {
		ticketDto.Assigned = uuid.Nil
	}

	return &Ticket{
		ID:        ticketDto.ID,
		RoomID:    ticketDto.RoomID,
		Metadata:  ticketDto.Metadata,
		IssuedBy:  ticketDto.IssuedBy,
		Assigned:  ticketDto.Assigned,
		Status:    ticketDto.Status,
		CreatedAt: ticketDto.CreatedAt,
		UpdatedAt: ticketDto.UpdatedAt,
	}

}
