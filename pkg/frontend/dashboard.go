package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/templates"
)

func viewDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Page("Template | Dashboard", templates.DashboardPage(),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}
