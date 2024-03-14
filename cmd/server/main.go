package main

import (
	"context"
	"fmt"
	"go-monolith-template/pkg/config"
	"go-monolith-template/pkg/core"
	"go-monolith-template/pkg/http_middleware"
	"go-monolith-template/pkg/logging"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/password_handling"
	"go-monolith-template/pkg/server"
	"go-monolith-template/pkg/session_handling"
	"go-monolith-template/pkg/store"
	"log/slog"
	_ "modernc.org/sqlite"
	"net/http"
	"strings"
	"time"
)

func main() {
	parsedConfig := config.Read()
	// Configure the global logger with the provided configuration options, sane defaults are already set to info level
	logging.ConfigureGlobalLoggerOptions(decodeLoggerValue(parsedConfig.Logging.Level), parsedConfig.Logging.AdditionalKeyPairs)
	// Wire database layer
	storageLayer := wireDatabase(parsedConfig.Database)
	// Ensure that there is at least one user in the database
	err := guardFromNoUsers(storageLayer)
	if err != nil {
		panic(err)
	}
	// Wire user management service
	userMgt := core.NewUserService(core.UserServiceOptions{
		StorageLayer:             storageLayer,
		LockoutThreshold:         parsedConfig.Security.FailedLoginLockoutAttempts,
		ComplexPasswordsRequired: parsedConfig.Security.RequireComplexPasswords,
		MinPasswordLength:        parsedConfig.Security.MinPasswordLength,
		MFAMandatory:             parsedConfig.Security.EnforceMFA,
		LockoutDurationMinutes:   parsedConfig.Security.FailedLoginLockoutTimeMinutes,
	})
	sessionHandler := session_handling.NewSessionManager(
		session_handling.NewSessionManagerOptions{
			MaxSessionAgeSeconds:   parsedConfig.Security.MaxSessionAgeSeconds,
			CookieName:             parsedConfig.Security.CookieName,
			CookieSecret:           parsedConfig.Security.CookieSecret,
			DefaultCacheExpiration: 30 * time.Second,
			SameSite:               http.SameSiteLaxMode,
			Secure:                 parsedConfig.Security.SecureCookie,
			HTTPOnly:               parsedConfig.Security.HTTPOnlyCookie,
			StorageLayer:           storageLayer,
		})
	middle := http_middleware.NewMiddleware(sessionHandler)
	s := server.NewServer(server.NewServerOptions{
		Port:                  parsedConfig.Server.Port,
		StorageLayer:          storageLayer,
		UserManagementService: userMgt,
		Middleware:            middle,
		SessionHandler:        sessionHandler,
	})
	defer storageLayer.Close()
	s.Run()
}

func decodeLoggerValue(level string) slog.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func wireDatabase(dbConfig config.Database) *store.Storage {
	switch strings.ToLower(dbConfig.Backend) {
	case "sqlite":
		s, err := store.ConnectSQLITE3(store.SqliteAdapterOptions{
			InMemory: false,
			FileName: dbConfig.File,
		}, true)
		if err != nil {
			panic(err)
		}
		return store.NewStorage(s)

	case "postgres":
		s, err := store.ConnectPostgres(store.PostgresAdapterOptions{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			Username: dbConfig.User,
			Password: dbConfig.Password,
			Database: dbConfig.Name,
			SSLMode:  dbConfig.SSLMode,
		}, true)
		if err != nil {
			panic(err)
		}
		return store.NewStorage(s)
	}
	panic(fmt.Sprintf("Unknown database backend: %v", dbConfig.Backend))
}

func guardFromNoUsers(s *store.Storage) error {
	users, err := s.UserList(context.TODO(), 100, 0)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		_, err := s.UserCreate(context.TODO(), store.CreateUserOptions{
			Email:        "admin@example.com",
			PasswordHash: password_handling.HashAndSaltPassword("Password1234!"),
			Role:         models.ADMIN_ROLE,
			MFASecret:    "",
			MFARequired:  false,
			APIKey:       password_handling.GenerateRandomPassword(32),
			Invited:      false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
