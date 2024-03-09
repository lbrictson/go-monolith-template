package frontend

import (
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/web"
	"io/fs"
	"net/http"
)

type LoadFrontendViewOptions struct {
	UserManagementService *core.UserService
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
	// Dashboard handlers
	e.GET("/", viewDashboard())
	return
}
