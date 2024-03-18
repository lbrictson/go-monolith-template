package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	Server   Server
	Database Database
	Email    Email
	Security Security
	Logging  Logging
}

type Server struct {
	Port        int    `default:"8080"`
	CallbackURL string `default:"http://localhost:8080"`
}

type Database struct {
	Backend  string `default:"sqlite"`
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Name     string `default:"postgres"`
	User     string `default:"postgres"`
	SSLMode  string `default:"disable"`
	Password string `default:"postgres"`
	File     string `default:"local/data/go-monolith-template.db"`
}

type Email struct {
	Mode     string `default:"smtp"`
	Host     string `default:"localhost"`
	Port     int    `default:"1025"`
	Username string `default:"admin"`
	Password string `default:"admin"`
	From     string `default:"no-reply@example.com"`
	SSLMode  string `default:"none"`
}

type Security struct {
	CookieSecret                  string `default:"thisshouldbechangedinprod"`
	SecureCookie                  bool   `default:"false"`
	HTTPOnlyCookie                bool   `default:"true"`
	CookieName                    string `default:"go-monolith-template"`
	MinPasswordLength             int    `default:"8"`
	RequireComplexPasswords       bool   `default:"true"`
	EnforceMFA                    bool   `default:"false"`
	FailedLoginLockout            bool   `default:"true"`
	FailedLoginLockoutTimeMinutes int    `default:"5"`
	FailedLoginLockoutAttempts    int    `default:"5"`
	MaxSessionAgeSeconds          int    `default:"86400"` // 24 hours
}

type Logging struct {
	Level              string            `default:"info"`
	AdditionalKeyPairs map[string]string `default:"appName:go-monolith-template,env:development"`
}

// Read reads the configuration from the environment.  All keys are read from their section with the prefix "app_$section"
// where $section is the name of the section in the struct.  For example, the Server.Port key is read from the
// environment variable "app_server_port".
func Read() *Config {
	server := Server{}
	err := envconfig.Process("app_server", &server)
	if err != nil {
		fmt.Printf("Fatal error reading server config: %v\n", err)
		os.Exit(1)
	}
	database := Database{}
	err = envconfig.Process("app_database", &database)
	if err != nil {
		fmt.Printf("Fatal error reading database config: %v\n", err)
		os.Exit(1)
	}
	email := Email{}
	err = envconfig.Process("app_email", &email)
	if err != nil {
		fmt.Printf("Fatal error reading email config: %v\n", err)
		os.Exit(1)
	}
	security := Security{}
	err = envconfig.Process("app_security", &security)
	if err != nil {
		fmt.Printf("Fatal error reading security config: %v\n", err)
		os.Exit(1)
	}
	logging := Logging{}
	err = envconfig.Process("app_logging", &logging)
	if err != nil {
		fmt.Printf("Fatal error reading logging config: %v\n", err)
		os.Exit(1)
	}
	c := Config{
		Server:   server,
		Database: database,
		Email:    email,
		Security: security,
		Logging:  logging,
	}
	return &c
}
