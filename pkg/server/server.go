package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/frontend"
	"go-monolith-template/pkg/http_middleware"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/pkg/store"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type NewServerOptions struct {
	Port                  int
	StorageLayer          *store.Storage
	UserManagementService *core.UserService
	Middleware            *http_middleware.Middleware
	SessionHandler        *session_handling.SessionManager
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		port:                  opts.Port,
		storageLayer:          opts.StorageLayer,
		userManagementService: opts.UserManagementService,
		middleware:            opts.Middleware,
		sessionHandler:        opts.SessionHandler,
	}
}

type Server struct {
	port                  int
	storageLayer          *store.Storage
	userManagementService *core.UserService
	middleware            *http_middleware.Middleware
	sessionHandler        *session_handling.SessionManager
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestID())
	frontend.LoadFrontendRoutes(e, frontend.LoadFrontendViewOptions{
		UserManagementService: s.userManagementService,
		SessionManager:        s.sessionHandler,
		MiddlewareManager:     s.middleware,
	})
	go func() {
		if err := e.Start(fmt.Sprintf(":%v", s.port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 3 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
