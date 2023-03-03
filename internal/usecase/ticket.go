package usecase

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
)

var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrUserNotFound   = errors.New("user not found")
	ErrRoomNotFound   = errors.New("room not found")
	ErrUnauthorized   = errors.New("unauthorized, user not in room")
)

func (u *Usecase) CreateTicket(ticket *entity.TicketDTO) error {
	if ticket == nil {
		return errors.New("nil ticket")
	}
	if err := ticket.Validate(); err != nil {
		return err
	}

	ok, err := u.UserRoom.IsUserInRoom(ticket.IssuedBy, ticket.RoomID)
	if err != nil {
		return err
	}

	if !ok {
		return ErrUnauthorized
	}

	return u.Ticket.CreateTicket(model.TicketToModel(u.uuidfn(), ticket))
}

func (u *Usecase) GetTicketByID(userID, id uuid.UUID) (*entity.TicketDTO, error) {
	if id == uuid.Nil {
		return nil, errors.New("id is empty")
	}

	return u.Ticket.GetTicketByID(userID, id)
}

func (u *Usecase) GetAllTickets(filter *entity.TicketFilter) ([]*entity.TicketDTO, int, error) {
	log.Info(filter)
	if filter == nil {
		return nil, 0, errors.New("filter is empty")
	}

	ok, err := u.UserRoom.IsUserInRoom(filter.IssuedBy, filter.RoomID)
	if err != nil {
		return nil, 0, err
	}

	if !ok {
		return nil, 0, ErrUnauthorized
	}

	if err = filter.Validate(); err != nil {
		return nil, 0, err
	}

	return u.Ticket.GetAllTickets(filter)
}

func (u *Usecase) UpdateTicket(userID uuid.UUID, ticket *entity.TicketDTO) error {
	if ticket == nil {
		return errors.New("ticket is empty")
	}

	oldTicket, err := u.Ticket.GetTicketByID(userID, ticket.ID)
	if err != nil {
		return err
	}
	oldTicket.Update(ticket)

	if err := oldTicket.Validate(); err != nil {
		return err
	}

	return u.Ticket.UpdateTicket(userID, model.TicketToModel(u.uuidfn(), oldTicket))
}

func (u *Usecase) DeleteTicket(userID uuid.UUID, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("id is empty")
	}

	return u.Ticket.DeleteTicket(userID, id)
}
