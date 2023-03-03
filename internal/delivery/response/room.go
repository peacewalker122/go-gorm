package response

import (
	"github.com/labstack/echo/v4"
	"go-gorm/entity"
)

func RoomResponse(c echo.Context, data *entity.RoomDTO) error {
	return c.JSON(200, echo.Map{
		"id":          data.ID,
		"name":        data.Name,
		"created_by":  data.CreatedBy,
		"description": data.Description,
		"max_members": data.MaxMembers,
		"created_at":  data.CreatedAt,
	})
}
