package core

import (
	"context"
	"errors"
	"fmt"
	"go-monolith-template/pkg/logging"
	"go-monolith-template/pkg/password_handling"
	"go-monolith-template/pkg/store"
	"strings"
	"time"
)

type UserService struct {
	dbConn                 *store.Storage
	lockoutTracker         map[string]int
	lockoutThreshold       int
	minPasswordLen         int
	complexPasswords       bool
	mfaMandatory           bool
	lockoutDurationMinutes int
}

type UserServiceOptions struct {
	StorageLayer             *store.Storage
	LockoutThreshold         int
	ComplexPasswordsRequired bool
	MinPasswordLength        int
	MFAMandatory             bool
	LockoutDurationMinutes   int
}

func NewUserService(opts UserServiceOptions) *UserService {
	return &UserService{
		dbConn:                 opts.StorageLayer,
		lockoutTracker:         make(map[string]int),
		lockoutThreshold:       opts.LockoutThreshold,
		minPasswordLen:         opts.MinPasswordLength,
		complexPasswords:       opts.ComplexPasswordsRequired,
		mfaMandatory:           opts.MFAMandatory,
		lockoutDurationMinutes: opts.LockoutDurationMinutes,
	}
}

// AuthenticateUser authenticates a user based on their email and password, all errors returned are frontend friendly
// the second bool returned indicates if MFA is required, if it is the user should be redirected to the MFA page
func (u *UserService) AuthenticateUser(ctx context.Context, email string, password string) (bool, bool, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("unknown user attempted to login", "email", email)
		return false, false, errors.New("Invalid username or password")
	}
	if user.Locked {
		fmt.Println("user was locked")
		fmt.Println(user.LockedAt)
		// If the lockout time has passed, unlock the user
		if !user.LockedAt.Add(time.Duration(u.lockoutDurationMinutes) * time.Minute).Before(time.Now()) {
			fmt.Println("unlocking user")
			f := false
			_, err := u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
				Locked:        &f,
				ClearLockedAt: true,
			})
			if err != nil {
				logging.FromContext(ctx).Error("failed to unlock user", "email", email, "error", err)
			}
		} else {
			logging.FromContext(ctx).Warn("locked user attempted to login", "email", email)
			return false, false, errors.New("Account is locked, please try again later")
		}
	}
	if password_handling.ComparePassword(user.PasswordHash, password) {
		// Regardless of MFA clear the lockout counter because of successful login
		delete(u.lockoutTracker, email)
		if user.MFAEnabled {
			return true, true, nil
		}
		now := time.Now()
		_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
			LastLogin: &now,
		})
		if err != nil {
			logging.FromContext(ctx).Error("failed to update last login", "email", email, "error", err)
		}
		return true, false, nil
	}
	u.lockoutTracker[email]++
	if u.lockoutTracker[email] >= u.lockoutThreshold {
		t := true
		now := time.Now()
		_, err := u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
			Locked:   &t,
			LockedAt: &now,
		})
		if err != nil {
			logging.FromContext(ctx).Error("failed to lock user", "email", email, "error", err)
		}
		// Clear the lockout counter because the user is now locked in the DB, no need to retain the counter for them
		delete(u.lockoutTracker, email)
		return false, false, errors.New("Account is locked, please try again in 5 minutes")
	}
	return false, false, errors.New("Invalid username or password")
}
