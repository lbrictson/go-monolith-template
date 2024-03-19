package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/templates"
)

func viewDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Page("Template | Dashboard", templates.DashboardPage(c.Get("email").(string), c.Get("isAdmin").(bool)),
			c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
	}
}
