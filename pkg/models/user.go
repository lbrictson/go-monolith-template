package models

import (
	"github.com/google/uuid"
	"time"
)

const (
	ADMIN_ROLE = "Admin"
	USER_ROLE  = "User"
)

type User struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"` // Ignored in JSON because it's sensitive
	Locked       bool       `json:"locked"`
	LockedAt     *time.Time `json:"locked_at"`
	APIKey       string     `json:"-"` // Ignored in JSON because it's sensitive
	MFAEnabled   bool       `json:"mfa_enabled"`
	MFASecret    string     `json:"-"` // Ignored in JSON because it's sensitive
	LastLogin    *time.Time `json:"last_login"`
	Invited      bool       `json:"invited"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Role         string     `json:"role"`
}