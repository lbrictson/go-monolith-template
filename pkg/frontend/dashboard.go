package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/models"
	"go-monolith-template/templates"
)

func viewDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)
		isAdmin := false
		if role == models.ADMIN_ROLE {
			isAdmin = true
		}
		return templates.Page("Template | Dashboard", templates.DashboardPage(c.Get("email").(string), isAdmin),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}
