package frontend

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/pkg/store"
	"go-monolith-template/templates"
	"strconv"
)

func viewTeam(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := usrService.GetUserByEmail(c.Request().Context(), c.Get("email").(string))
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil).Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.Page("Template | Team", templates.PageAdminTeam(*u, c.Get("isAdmin").(bool)),
			c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
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
		return templates.TeamTable(u, self, page, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxCreateUserForm(u *core.UserService, sess *session_handling.SessionManager) echo.HandlerFunc {
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
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: err.Error(),
						IsError: false,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		users, err := u.ListUsers(ctx, 0, 10)
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: err.Error(),
						IsError: false,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.TeamTable(users, c.Get("email").(string),
			0, []session_handling.Notification{
				{
					Header:  "Success",
					Message: fmt.Sprintf("%v created", form.Email),
					IsError: false,
				}}).Render(c.Request().Context(), c.Response().Writer)
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
			0, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}

func formAdminSetPassword(u *core.UserService, sess *session_handling.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		type Form struct {
			Email    string `form:"email" validate:"required,email"`
			Password string `form:"password" validate:"required"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil).Render(c.Request().Context(), c.Response().Writer)
		}
		err := u.SetUserPassword(ctx, form.Email, form.Password)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil).Render(c.Request().Context(), c.Response().Writer)
		}
		sess.AddNotificationViaContext(c, "Success", fmt.Sprintf("%v password updated", form.Email), false)
		return c.Redirect(302, "/admin/team")
	}
}

func viewAdminSetPassword(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		targetUserID := c.Param("id")
		user, err := u.GetUserByID(ctx, uuid.MustParse(targetUserID))
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil).Render(ctx, c.Response().Writer)
		}
		return templates.Page("Template | Team", templates.AdminSetPasswordPage(c.Get("email").(string), c.Get("isAdmin").(bool), *user),
			c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxAdminSwapRole(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, _ := u.ListUsers(ctx, 0, 10)
		id := c.Param("id")
		user, err := u.GetUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		if user.Role == models.ADMIN_ROLE {
			user.Role = models.USER_ROLE
		} else {
			user.Role = models.ADMIN_ROLE
		}
		err = u.UpdateUser(ctx, uuid.MustParse(id), store.UpdateUserOptions{
			Role: &user.Role,
		})
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		users, _ = u.ListUsers(ctx, 0, 10)
		return templates.TeamTable(users, c.Get("email").(string),
			0, []session_handling.Notification{
				{
					Header:  "Success",
					Message: fmt.Sprintf("%v role changed to %v", user.Email, user.Role),
					IsError: false,
				}}).Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxAdminDisableMFA(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, _ := u.ListUsers(ctx, 0, 10)
		id := c.Param("id")
		user, err := u.GetUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		err = u.DisableMFA(ctx, user.Email)
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		users, _ = u.ListUsers(ctx, 0, 10)
		return templates.TeamTable(users, c.Get("email").(string),
			0, []session_handling.Notification{
				{
					Header:  "Success",
					Message: fmt.Sprintf("%v MFA disabled", user.Email),
					IsError: false,
				}}).Render(c.Request().Context(), c.Response().Writer)
	}
}

func htmxAdminDeleteUser(u *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, _ := u.ListUsers(ctx, 0, 10)
		id := c.Param("id")
		user, err := u.GetUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		err = u.DeleteUser(ctx, user.Email)
		if err != nil {
			return templates.TeamTable(users, c.Get("email").(string),
				0, []session_handling.Notification{
					{
						Header:  "Error",
						Message: fmt.Sprintf("%v", err.Error()),
						IsError: true,
					}}).Render(c.Request().Context(), c.Response().Writer)
		}
		users, _ = u.ListUsers(ctx, 0, 10)
		return templates.TeamTable(users, c.Get("email").(string),
			0, []session_handling.Notification{
				{
					Header:  "Success",
					Message: fmt.Sprintf("%v deleted", user.Email),
					IsError: false,
				}}).Render(c.Request().Context(), c.Response().Writer)
	}
}
