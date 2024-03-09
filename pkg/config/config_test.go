package config

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	// Validate default values are set
	defaults := Config{
		Server: Server{
			Port:        8080,
			CallbackURL: "http://localhost:8080",
		},
		Database: Database{
			Backend:  "sqlite",
			Host:     "localhost",
			Port:     5432,
			Name:     "postgres",
			User:     "postgres",
			SSLMode:  "disable",
			Password: "postgres",
			File:     "local/data/go-monolith-template.db",
		},
		Email: Email{
			Mode:     "smtp",
			Host:     "localhost",
			Port:     1025,
			Username: "admin",
			Password: "admin",
			From:     "no-reply@example.com",
			SSLMode:  "none",
		},
		Security: Security{
			CookieSecret:                  "thisshouldbechangedinprod",
			SecureCookie:                  false,
			HTTPOnlyCookie:                true,
			CookieName:                    "go-monolith-template",
			MinPasswordLength:             8,
			RequireComplexPasswords:       true,
			EnforceMFA:                    false,
			FailedLoginLockout:            true,
			FailedLoginLockoutTimeMinutes: 5,
			FailedLoginLockoutAttempts:    5,
		},
		Logging: Logging{
			Level: "info",
			AdditionalKeyPairs: map[string]string{
				"appName": "go-monolith-template",
				"env":     "development",
			},
		},
	}
	c := Read()
	if !reflect.DeepEqual(c.Database, defaults.Database) {
		t.Errorf("Expected default values: \n%+v\ngot: \n%+v", defaults.Database, c.Database)
	}
	if !reflect.DeepEqual(c.Server, defaults.Server) {
		t.Errorf("Expected default values: \n%+v\ngot: \n%+v", defaults.Server, c.Server)
	}
	if !reflect.DeepEqual(c.Email, defaults.Email) {
		t.Errorf("Expected default values: \n%+v\ngot: \n%+v", defaults.Email, c.Email)
	}
	if !reflect.DeepEqual(c.Security, defaults.Security) {
		t.Errorf("Expected default values: \n%+v\ngot: \n%+v", defaults.Security, c.Security)
	}
	if !reflect.DeepEqual(c.Logging, defaults.Logging) {
		t.Errorf("Expected default values: \n%+v\ngot: \n%+v", defaults.Logging, c.Logging)
	}
	// Set a few environment variables and validate they are set
	_ = os.Setenv("APP_SERVER_PORT", "8081")
	_ = os.Setenv("APP_DATABASE_BACKEND", "postgres")
	_ = os.Setenv("APP_SECURITY_SECURECOOKIE", "true")
	_ = os.Setenv("APP_EMAIL_MODE", "ses")
	_ = os.Setenv("APP_LOGGING_LEVEL", "debug")
	_ = os.Setenv("APP_LOGGING_ADDITIONALKEYPAIRS", "env:production,logType:file")
	c = Read()
	if c.Server.Port != 8081 {
		t.Errorf("Expected Server.Port to be 8081, got %d", c.Server.Port)
	}
	if c.Database.Backend != "postgres" {
		t.Errorf("Expected Database.Backend to be postgres, got %s", c.Database.Backend)
	}
	if !c.Security.SecureCookie {
		t.Errorf("Expected Security.SecureCookie to be true, got false")
	}
	if c.Email.Mode != "ses" {
		t.Errorf("Expected Email.Mode to be ses, got %s", c.Email.Mode)
	}
	if c.Logging.Level != "debug" {
		t.Errorf("Expected Logging.Level to be debug, got %s", c.Logging.Level)
	}
	if c.Logging.AdditionalKeyPairs["env"] != "production" {
		t.Errorf("Expected Logging.AdditionalKeyPairs[env] to be production, got %s", c.Logging.AdditionalKeyPairs["env"])
	}
	if c.Logging.AdditionalKeyPairs["logType"] != "file" {
		t.Errorf("Expected Logging.AdditionalKeyPairs[logType] to be file, got %s", c.Logging.AdditionalKeyPairs["logType"])
	}
}
