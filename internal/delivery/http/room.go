package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go-gorm/internal/delivery/request"
	"go-gorm/internal/delivery/response"
	"go-gorm/pkg/middleware"
)

func (h *Handler) CreateRoom(c echo.Context) error {
	req := new(request.RoomRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(500, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}
	payload := middleware.GetAuthPayload(c)

	roomDto := req.ToDto()
	roomDto.CreatedBy = uuid.Must(uuid.Parse(payload.Id))

	err := h.RoomUsecase.CreateRoom(c.Request().Context(), roomDto)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return response.SuccessResponse(c)
}

func (h *Handler) GetRoom(c echo.Context) error {
	req := new(request.RoomQueryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(500, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	h.echo.Logger.Infof("from struct: %s", req.ID)
	log.Info("from struct: ", req.ID)
	roomId := c.Param("id")

	h.echo.Logger.Info("from param: ", roomId)
	log.Info("from param: ", roomId)
	if req.ID == "" {
		req.ID = roomId
	}
	payload := middleware.GetAuthPayload(c)
	roomuuid := uuid.Must(uuid.Parse(req.ID))

	if payload.GetID() == uuid.Nil || roomuuid == uuid.Nil {
		return c.JSON(400, "invalid id")
	}

	room, err := h.RoomUsecase.GetRoomByID(c.Request().Context(), roomuuid, payload.GetID())
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return response.RoomResponse(c, room)
}
