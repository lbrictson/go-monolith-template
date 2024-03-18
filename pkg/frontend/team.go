package frontend

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/store"
	"go-monolith-template/templates"
	"strconv"
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

func htmxViewTeamTable(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		self := c.Get("email").(string)
		limit := 10
		page := 0
		if c.QueryParam("limit") != "" {
			parsedLimit, err := strconv.Atoi(c.QueryParam("limit"))
			if err != nil {
				return c.JSON(400, map[string]string{"error": "Invalid limit"})
			}
			limit = parsedLimit
		}
		if c.QueryParam("page") != "" {
			parsedPage, err := strconv.Atoi(c.QueryParam("page"))
			if err != nil {
				return c.JSON(400, map[string]string{"error": "Invalid page"})
			}
			page = parsedPage
		}

		u, err := usrService.ListUsers(c.Request().Context(), page, limit)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return templates.TeamTable(u, self, page, false, "", false, "").Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxCreateUserForm(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		type Form struct {
			Email    string `form:"email" validate:"required,email"`
			Password string `form:"password" validate:"required"`
			Role     string `form:"role" validate:"required"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid form"})
		}
		_, err := u.CreateUser(ctx, form.Email, form.Role, form.Password)
		if err != nil {
			users, _ := u.ListUsers(ctx, 0, 10)
			return templates.TeamTable(users, c.Get("email").(string),
				0, true, err.Error(), false, "").Render(c.Request().Context(), c.Response().Writer)
		}
		users, err := u.ListUsers(ctx, 0, 10)
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, true, err.Error(), false, "").Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.TeamTable(users, c.Get("email").(string),
			0, false, "", true, fmt.Sprintf("%v created", form.Email)).Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxTeamSearchForm(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			Email string `form:"email"`
			Role  string `form:"role"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid form"})
		}
		opts := store.UserQueryOptions{}
		if form.Email != "" {
			opts.EmailLike = &form.Email
		}
		if form.Role != "" {
			opts.RoleIs = &form.Role
		}
		users, err := u.QueryUsers(c.Request().Context(), opts)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return templates.TeamTable(users, c.Get("email").(string),
			0, false, "", false, "").Render(c.Request().Context(), c.Response().Writer)
	}
}
