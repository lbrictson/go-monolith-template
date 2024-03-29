package frontend

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/templates"
	"image/png"
	"net/http"
)

func viewProfile(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := usrService.GetUserByEmail(c.Request().Context(), c.Get("email").(string))
		if err != nil {
			return templates.ErrorPage(err.Error()).Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.Page("Template | Profile", templates.ProfilePage(*u, c.Get("isAdmin").(bool)),
			c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
	}
}

func viewEnableMFA(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := usrService.GetUserByEmail(c.Request().Context(), c.Get("email").(string))
		if err != nil {
			return templates.ErrorPage(err.Error()).Render(c.Request().Context(), c.Response().Writer)
		}
		if u.MFAEnabled {
			return templates.ErrorPage("MFA is already enabled").Render(c.Request().Context(), c.Response().Writer)
		}
		var buf bytes.Buffer
		secret, img, err := usrService.GenerateMFAAssets(c.Request().Context(), u.Email)
		png.Encode(&buf, img)
		imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
		secretForImageBlock := fmt.Sprintf("data:image/png;base64,%s", imgBase64Str)
		return templates.Page("Template | Enable MFA", templates.PageEnableMFA(*u, c.Get("isAdmin").(bool), secretForImageBlock, secret),
			c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
	}
}

func formSubmitEnableMFA(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Get("email").(string)
		token := c.FormValue("token")
		err := usrService.EnableMFA(c.Request().Context(), email, token)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
		}
		return c.Redirect(http.StatusFound, "/profile")
	}
}

func formSubmitUpdatePassword(usrService *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Get("email").(string)
		newPassword := c.FormValue("password")
		err := usrService.UpdatePassword(c.Request().Context(), email, newPassword)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage("Invalid password"),
				c.Get("notifications").([]session_handling.Notification)).Render(c.Request().Context(), c.Response().Writer)
		}
		return c.Redirect(http.StatusFound, "/profile")
	}
}
