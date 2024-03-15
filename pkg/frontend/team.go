package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/models"
	"go-monolith-template/templates"
)

func viewTeam(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := false
		if c.Get("role").(string) == models.ADMIN_ROLE {
			isAdmin = true
		}
		u, err := usrService.GetUserByEmail(c.Request().Context(), c.Get("email").(string))
		if err != nil {
			return templates.ErrorPage(err.Error()).Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.Page("Template | Team", templates.PageAdminTeam(*u, isAdmin),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}
