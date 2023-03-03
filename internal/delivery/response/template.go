package response

import "github.com/labstack/echo/v4"

func SuccessResponse(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "Success",
	})
}

func GetAllResponse(c echo.Context, data interface{}, total, page, limit int) error {
	return c.JSON(200, getallresponse{
		Message: "Success",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

type getallresponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	Limit   int         `json:"per-page"`
}
