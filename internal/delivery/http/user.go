package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-gorm/internal/delivery/request"
	"go-gorm/internal/delivery/response"
)

func (h *Handler) GetUserProfile(c echo.Context) error {
	req := new(request.QueryUserRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return c.JSON(400, err)
	}

	user, err := h.UserUsecase.GetUserByID(id)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	req := new(request.UserRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	reqDto := req.ToUserDto()

	err := h.UserUsecase.CreateUser(reqDto)
	if err != nil {
		return c.JSON(400, err)
	}

	return response.SuccessResponse(c)
}

func (h *Handler) Login(c echo.Context) error {
	req := new(request.LoginRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	user, token, err := h.UserUsecase.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(err.ResponseError())
	}

	return response.LoginResponse(c, token, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	req := new(request.UserRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	reqDto := req.ToUserDto()

	err := h.UserUsecase.UpdateUser(reqDto)
	if err != nil {
		return c.JSON(400, err)
	}

	return response.SuccessResponse(c)
}
