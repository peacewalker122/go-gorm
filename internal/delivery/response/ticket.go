package response

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-gorm/entity"
	"time"
)

func TicketResponse(c echo.Context, ticket *entity.TicketDTO) error {
	return c.JSON(200, ticketresponse{
		Id:        ticket.ID,
		Room:      ticket.RoomID,
		IssuedBy:  ticket.IssuedBy,
		Assigned:  ticket.Assigned,
		Subject:   ticket.Metadata,
		Status:    ticket.Status.String(),
		CreatedAt: ticket.CreatedAt.Format(time.RFC850),
		UpdatedAt: ticket.UpdatedAt.Format(time.RFC850),
	})
}

type ticketresponse struct {
	Id        uuid.UUID `json:"id"`
	Room      uuid.UUID `json:"room"`
	IssuedBy  uuid.UUID `json:"issuedBy"`
	Assigned  uuid.UUID `json:"assigned"`
	Subject   string    `json:"subject"`
	Status    string    `json:"status"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

func TicketsDtoToResponse(ticket []*entity.TicketDTO) []*ticketresponse {
	var tickets []*ticketresponse
	for _, t := range ticket {
		tickets = append(tickets, &ticketresponse{
			Id:        t.ID,
			Room:      t.RoomID,
			IssuedBy:  t.IssuedBy,
			Assigned:  t.Assigned,
			Subject:   t.Metadata,
			Status:    t.Status.String(),
			CreatedAt: t.CreatedAt.Format(time.RFC850),
			UpdatedAt: t.UpdatedAt.Format(time.RFC850),
		})
	}

	return tickets
}
