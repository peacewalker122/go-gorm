package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-gorm/internal/delivery/request"
	"go-gorm/internal/delivery/response"
	"go-gorm/pkg/middleware"
)

func (h *Handler) CreateTicket(c echo.Context) error {
	req := new(request.TicketRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(500, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}
	payload := middleware.GetAuthPayload(c)

	ticketDto := req.ToDto()
	ticketDto.IssuedBy = uuid.Must(uuid.Parse(payload.Id))

	err := h.TicketUsecase.CreateTicket(ticketDto)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return response.SuccessResponse(c)
}

func (h *Handler) GetTicketByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, "id is empty")
	}

	payload := middleware.GetAuthPayload(c)

	ticket, err := h.TicketUsecase.GetTicketByID(uuid.Must(uuid.Parse(payload.Id)), uuid.Must(uuid.Parse(id)))
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return response.TicketResponse(c, ticket)
}

func (h *Handler) GetAllTickets(c echo.Context) error {
	req := new(request.TicketFilterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(500, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	payload := middleware.GetAuthPayload(c)

	tickets, total, err := h.TicketUsecase.GetAllTickets(req.ToTicketFilter(uuid.Must(uuid.Parse(payload.Id))))
	if err != nil {
		return c.JSON(500, err.Error())
	}

	ticketResponse := response.TicketsDtoToResponse(tickets)

	return response.GetAllResponse(c, ticketResponse, total, req.Page, req.PageSize)
}
