package middleware

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type customValidator struct {
	validate *validator.Validate
}

func NewValidator(arg *validator.Validate) *customValidator {
	return &customValidator{
		validate: arg,
	}
}

func (v *customValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func HTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errmsg := []string{}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, e := range castedObject {
			errmsg = append(errmsg, fmt.Sprintf("error happen in %s, due %s", e.Field(), e.ActualTag()))
		}
	}

	c.JSON(report.Code, errmsg)
	makeLogEntry(c).Error(report.Message)
}
