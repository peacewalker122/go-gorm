package repository

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
	"go-gorm/pkg/enum"
)

type Ticket interface {
	CreateTicket(ticket *model.Ticket) error
	GetTicketByID(userID, id uuid.UUID) (*entity.TicketDTO, error)
	GetAllTickets(filter *entity.TicketFilter, opt ...string) ([]*entity.TicketDTO, int, error)
	UpdateTicket(userID uuid.UUID, ticket *model.Ticket) error
	UpdateTicketStatus(id, userID uuid.UUID, status enum.TicketStatus) error
	DeleteTicket(id, userID uuid.UUID) error
}
