package repository

import (
	"errors"
	"github.com/google/uuid"
	"go-gorm/domain/repository"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
	"go-gorm/pkg/enum"
	"gorm.io/gorm"
	"time"
)

type TicketRepository struct {
	db *gorm.DB
}

func (t *TicketRepository) CreateTicket(ticket *model.Ticket) error {
	err := t.db.Create(ticket).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepository) GetTicketByID(userID uuid.UUID, id uuid.UUID) (*entity.TicketDTO, error) {
	var (
		ticketModel = &model.Ticket{}
		whereQuery  = t.db
		joinQuery   = t.db
	)

	err := t.db.First(ticketModel, id).Error
	if err != nil {
		return nil, err
	}

	joinQuery = joinQuery.
		InnerJoins("inner join user_room ur on ur.room_id = ticket.room_id").
		InnerJoins("inner join \"user\" u on u.id = ur.user_id").Where(
		whereQuery.
			Where("u.id = ?", userID).
			Where("ticket.room_id = ?", ticketModel.RoomID).
			Where("u.deleted_at is null"),
	).Select("ticket.*").Model(&model.Ticket{})

	err = joinQuery.Scan(ticketModel).Error

	if ticketModel.ID == uuid.Nil {
		return nil, errors.New("ticket not found")
	}

	if err != nil {
		return nil, err
	}

	return ticketModel.ToDTO(), nil
}

func (t *TicketRepository) GetAllTickets(filter *entity.TicketFilter, opt ...string) ([]*entity.TicketDTO, int, error) {
	var (
		ticketModels []*model.Ticket
		subquery     = t.db
		whereQuery   = t.db
	)

	offset := (filter.Page - 1) * filter.PageSize

	dbResult := t.db.Model(&ticketModels).
		Where(whereQuery.
			Where("status = ?", filter.Status.String()).
			Where("room_id = ?", filter.RoomID),
		)

	switch {
	case filter.Assigned != uuid.Nil:
		subquery = subquery.Where("assigned = ?", filter.Assigned)
	case filter.IssuedBy != uuid.Nil:
		subquery = subquery.Where("issued_by = ?", filter.IssuedBy)
	case filter.TimeFilter != time.Time{}:
		if filter.TimeFilter.After(time.Now()) {
			subquery = subquery.Order("created_at asc")
		} else {
			subquery = subquery.Order("created_at desc")
		}
	}

	dbResult = dbResult.Where(subquery).
		Offset(offset).
		Limit(filter.PageSize).
		Find(&ticketModels)

	if dbResult.Error != nil {
		return nil, 0, dbResult.Error
	}
	var tickets []*entity.TicketDTO
	for _, ticket := range ticketModels {
		tickets = append(tickets, ticket.ToDTO())
	}
	return tickets, int(dbResult.RowsAffected), nil
}

func (t *TicketRepository) UpdateTicket(userID uuid.UUID, ticket *model.Ticket) error {
	err := t.db.Model(model.Ticket{}).Save(ticket).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepository) UpdateTicketStatus(userID, id uuid.UUID, status enum.TicketStatus) error {
	err := t.db.Model(&model.Ticket{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepository) DeleteTicket(id, userID uuid.UUID) error {
	err := t.db.Model(&model.Ticket{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("now()")).Error
	if err != nil {
		return err
	}

	return nil
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

var _ repository.Ticket = (*TicketRepository)(nil)
