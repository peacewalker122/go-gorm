package response

import (
	"github.com/labstack/echo/v4"
	"go-gorm/entity"
)

func UserResponse(c echo.Context, user *entity.UserDTO) error {
	return c.JSON(200, echo.Map{
		"id":        user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"status":    user.Status,
		"createdAt": user.CreatedAt,
	})
}

func LoginResponse(c echo.Context, tokenMap map[string]interface{}, user *entity.UserDTO) error {
	return c.JSON(200, echo.Map{
		"token": tokenMap,
		"user": echo.Map{
			"id":        user.ID,
			"email":     user.Email,
			"username":  user.Username,
			"status":    user.Status,
			"createdAt": user.CreatedAt,
		},
	})
}
