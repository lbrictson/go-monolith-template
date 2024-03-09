package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/templates"
)

func viewLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Page("Template | Login", templates.LoginPage(),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}
