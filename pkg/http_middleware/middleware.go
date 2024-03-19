package http_middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/logging"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/pkg/store"
)

type Middleware struct {
	sessionHandler SessionHandler
}

type SessionHandler interface {
	HasSession(c echo.Context) bool
	Save(c echo.Context, user models.User) error
	Destroy(c echo.Context) error
	GetSessionData(c echo.Context) (models.SessionData, error)
	Update(c echo.Context, opts store.UpdateSessionOptions) error
	GetNotifications(id uuid.UUID) []session_handling.Notification
}

func NewMiddleware(sessionHandler SessionHandler) *Middleware {
	return &Middleware{
		sessionHandler: sessionHandler,
	}
}

func (m *Middleware) LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !m.sessionHandler.HasSession(c) {
			return c.Redirect(302, "/login")
		}
		data, err := m.sessionHandler.GetSessionData(c)
		if err != nil {
			return c.Redirect(302, "/login")
		}
		c.Set("session_id", data.ID)
		c.Set("user_id", data.UserID)
		c.Set("email", data.Email)
		c.Set("role", data.Role)
		c.Set("mfa_enabled", data.MFAEnabled)
		c.Set("mfa_completed", data.MFACompleted)
		c.Set("auth_method", "frontend")
		c.Set("request_id", c.Response().Header().Get(echo.HeaderXRequestID))
		c.Set("ip", c.RealIP())
		if data.Role == models.ADMIN_ROLE {
			c.Set("isAdmin", true)
		} else {
			c.Set("isAdmin", false)
		}
		c.Set("notifications", m.sessionHandler.GetNotifications(data.ID))
		l := logging.FromEchoContext(c)
		l = l.With("http_session_id", c.Get("session_id"))
		l = l.With("http_user_id", c.Get("user_id"))
		l = l.With("http_email", c.Get("email"))
		l = l.With("http_role", c.Get("role"))
		l = l.With("http_auth_method", c.Get("auth_method"))
		l = l.With("http_ip", c.Get("ip"))
		l = l.With("http_method", c.Request().Method)
		l = l.With("http_request_id", c.Response().Header().Get(echo.HeaderXRequestID))
		c.Set("logger", l)
		c.SetRequest(c.Request().WithContext(logging.IntoContext(c.Request().Context(), l)))
		if data.MFAEnabled && !data.MFACompleted {
			return c.Redirect(302, "/mfa")
		}
		return next(c)
	}
}

func (m *Middleware) AdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)
		if role != models.ADMIN_ROLE {
			return c.JSON(403, "Forbidden")
		}
		return next(c)
	}
}
