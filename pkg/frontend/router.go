package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/http_middleware"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/web"
	"io/fs"
	"net/http"
)

type LoadFrontendViewOptions struct {
	UserManagementService *core.UserService
	SessionManager        *session_handling.SessionManager
	MiddlewareManager     *http_middleware.Middleware
}

func LoadFrontendRoutes(e *echo.Echo, options LoadFrontendViewOptions) {
	// Read in static assets from the mock file system
	fSys, err := fs.Sub(web.Assets, "static")
	if err != nil {
		panic(err)
	}
	assetHandler := http.FileServer(http.FS(fSys))
	// Static assets
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	// Auth handlers
	e.GET("/login", viewLogin())
	e.POST("/login", formSubmitLogin(options.UserManagementService, options.SessionManager))
	e.GET("/logout", hookLogout(options.SessionManager))
	e.GET("/mfa", viewMFA())
	e.POST("/mfa", formSubmitMFA(options.UserManagementService, options.SessionManager))
	// Dashboard handlers
	e.GET("/", viewDashboard(), options.MiddlewareManager.LoginRequired)
	// Profile handlers
	e.GET("/profile", viewProfile(options.UserManagementService), options.MiddlewareManager.LoginRequired)
	e.GET("/profile/enable_mfa", viewEnableMFA(options.UserManagementService), options.MiddlewareManager.LoginRequired)
	e.POST("/profile/enable_mfa", formSubmitEnableMFA(options.UserManagementService), options.MiddlewareManager.LoginRequired)
	e.DELETE("/profile/disable_mfa", hookDisableMFA(options.UserManagementService), options.MiddlewareManager.LoginRequired)
	e.POST("/profile/password", formSubmitUpdatePassword(options.UserManagementService), options.MiddlewareManager.LoginRequired)
	// Admin - Team handlers
	e.GET("/admin/team", viewTeam(options.UserManagementService), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	e.GET("/component/admin/team_table", htmxViewTeamTable(options.UserManagementService), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	e.POST("/component/admin/create_user", htmxCreateUserForm(options.UserManagementService, options.SessionManager), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	e.POST("/component/admin/search_user", htmxTeamSearchForm(options.UserManagementService), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	e.POST("/admin/team/set_password", formAdminSetPassword(options.UserManagementService, options.SessionManager), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	e.GET("/admin/team/set_password/:id", viewAdminSetPassword(options.UserManagementService), options.MiddlewareManager.LoginRequired, options.MiddlewareManager.AdminRequired)
	return
}
