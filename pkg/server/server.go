package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/frontend"
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
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		port:                  opts.Port,
		storageLayer:          opts.StorageLayer,
		userManagementService: opts.UserManagementService,
	}
}

type Server struct {
	port                  int
	storageLayer          *store.Storage
	userManagementService *core.UserService
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	e := echo.New()
	e.HideBanner = true
	frontend.LoadFrontendRoutes(e, frontend.LoadFrontendViewOptions{
		UserManagementService: s.userManagementService,
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
