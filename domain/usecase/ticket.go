package usecase

import (
	"github.com/google/uuid"
	"go-gorm/entity"
)

type TicketUsecase interface {
	CreateTicket(ticket *entity.TicketDTO) error
	GetTicketByID(userID, id uuid.UUID) (*entity.TicketDTO, error)
	GetAllTickets(filter *entity.TicketFilter) ([]*entity.TicketDTO, int, error)
	UpdateTicket(userID uuid.UUID, ticket *entity.TicketDTO) error
	DeleteTicket(userID, id uuid.UUID) error
}
