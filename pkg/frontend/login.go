package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/pkg/store"
	"go-monolith-template/templates"
	"net/http"
)

func viewLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Page("Template | Login", templates.LoginPage(),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}

func formSubmitLogin(a *core.UserService, s *session_handling.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		type form struct {
			Email    string `form:"email"`
			Password string `form:"password"`
		}
		f := new(form)
		if err := c.Bind(f); err != nil {
			return templates.Page("Template | Error", templates.ErrorPage("Error parsing form"),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		_, mfaNeeded, err := a.AuthenticateUser(c.Request().Context(), f.Email, f.Password)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		// Valid login credentials
		userInformation, err := a.GetUserByEmail(c.Request().Context(), f.Email)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		err = s.Save(c, *userInformation)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		if mfaNeeded {
			return c.Redirect(http.StatusFound, "/mfa?email="+f.Email)
		}
		return c.Redirect(http.StatusFound, "/")
	}
}

func hookLogout(s *session_handling.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := s.Destroy(c)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		return c.Redirect(http.StatusFound, "/login")
	}
}

func viewMFA() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.QueryParam("email")
		return templates.Page("Template | MFA", templates.MFAPage(email),
			nil, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}

func formSubmitMFA(a *core.UserService, s *session_handling.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		type form struct {
			Token string `form:"token"`
			Email string `form:"email"`
		}
		f := new(form)
		if err := c.Bind(f); err != nil {
			return templates.Page("Template | Error", templates.ErrorPage("Error parsing form"),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		valid := a.ValidateMFAToken(c.Request().Context(), f.Email, f.Token)
		if !valid {
			return templates.Page("Template | Error", templates.ErrorPage("Invalid MFA token"),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		t := true
		s.Update(c, store.UpdateSessionOptions{MFACompleted: &t})
		return c.Redirect(http.StatusFound, "/")
	}
}

func hookDisableMFA(a *core.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Get("email").(string)
		err := a.DisableMFA(c.Request().Context(), email)
		if err != nil {
			return templates.Page("Template | Error", templates.ErrorPage(err.Error()),
				nil, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		c.Response().Header().Set("HX-Redirect", "/profile")
		return c.String(http.StatusOK, "MFA has been disabled")
	}
}
