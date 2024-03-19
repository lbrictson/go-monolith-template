package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"go-monolith-template/mailer"
	"go-monolith-template/pkg/logging"
	"go-monolith-template/pkg/models"
	"go-monolith-template/pkg/password_handling"
	"go-monolith-template/pkg/store"
	"image"
	"strings"
	"time"
)

type UserService struct {
	dbConn                 *store.Storage
	emailService           mailer.Emailer
	serverURL              string
	lockoutTracker         map[string]int
	lockoutThreshold       int
	minPasswordLen         int
	complexPasswords       bool
	mfaMandatory           bool
	lockoutDurationMinutes int
}

type UserServiceOptions struct {
	StorageLayer             *store.Storage
	EmailService             mailer.Emailer
	ServerURL                string
	LockoutThreshold         int
	ComplexPasswordsRequired bool
	MinPasswordLength        int
	MFAMandatory             bool
	LockoutDurationMinutes   int
}

func NewUserService(opts UserServiceOptions) *UserService {
	return &UserService{
		dbConn:                 opts.StorageLayer,
		emailService:           opts.EmailService,
		serverURL:              opts.ServerURL,
		lockoutTracker:         make(map[string]int),
		lockoutThreshold:       opts.LockoutThreshold,
		minPasswordLen:         opts.MinPasswordLength,
		complexPasswords:       opts.ComplexPasswordsRequired,
		mfaMandatory:           opts.MFAMandatory,
		lockoutDurationMinutes: opts.LockoutDurationMinutes,
	}
}

func (u *UserService) SetPasswordViaResetToken(ctx context.Context, email string, token string, newPassword string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	if user.PasswordResetToken != token {
		logging.FromContext(ctx).Warn("invalid password reset token", "email", email)
		return errors.New("Invalid password reset token")
	}
	if user.PasswordResetTokenExpiration.Before(time.Now()) {
		logging.FromContext(ctx).Warn("expired password reset token", "email", email)
		return errors.New("Expired password reset token")
	}
	if !password_handling.IsPasswordValid(newPassword, u.minPasswordLen, u.complexPasswords) {
		return errors.New("Invalid password")
	}
	hash := password_handling.HashAndSaltPassword(newPassword)
	empty := ""
	_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
		PasswordHash:                      &hash,
		PasswordResetToken:                &empty,
		ClearPasswordResetTokenExpiration: true,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to update password", "email", email, "error", err)
		return errors.New("Failed to update password")
	}
	return nil
}

func (u *UserService) RequestPasswordReset(ctx context.Context, email string) {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return
	}
	token := uuid.New().String()
	exp := time.Now().Add(time.Hour * 24)
	_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
		PasswordResetToken:           &token,
		PasswordResetTokenExpiration: &exp,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to update password reset token", "email", email, "error", err)
		return
	}
	err = u.emailService.SendEmail(ctx, mailer.SendEmailInput{
		To:               user.Email,
		Subject:          "Reset your Go-Mono-Template password",
		CC:               "",
		BCC:              "",
		Title:            "Reset your Go-Mono-Template password",
		PoweredByLink:    u.serverURL,
		PoweredByText:    "Go Monolith Template",
		ContentHeader:    "Reset Password",
		ContentText:      "Click the button below to reset your password",
		CallToActionLink: fmt.Sprintf("%s/reset_password?token=%s&email=%s", u.serverURL, token, user.Email),
		CallToActionText: "Reset Password",
		UnsubscribeLink:  u.serverURL + "/unsubscribe",
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to send password reset email", "email", email, "error", err)
	}
}

func (u *UserService) DeleteUser(ctx context.Context, email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	err = u.dbConn.UserDelete(ctx, user.ID)
	if err != nil {
		logging.FromContext(ctx).Error("failed to delete user", "email", email, "error", err)
		return errors.New("Failed to delete user")
	}
	return nil
}

func (u *UserService) SetUserPassword(ctx context.Context, email string, newPassword string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	if !password_handling.IsPasswordValid(newPassword, u.minPasswordLen, u.complexPasswords) {
		return errors.New("Invalid password - does not meet complexity requirements")
	}
	hash := password_handling.HashAndSaltPassword(newPassword)
	_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
		PasswordHash: &hash,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to update password", "email", email, "error", err)
		return errors.New("Failed to update password")
	}
	return nil
}

func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := u.dbConn.UserGetByID(ctx, id)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "id", id)
		return nil, errors.New("User not found")
	}
	return user, nil
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
		// If the lockout time has passed, unlock the user
		if !user.LockedAt.Add(time.Duration(u.lockoutDurationMinutes) * time.Minute).Before(time.Now()) {
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
		f := false
		_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
			LastLogin: &now,
			Invited:   &f, // Remove invited flag because the user has now logged in
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

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (u *UserService) EnableMFA(ctx context.Context, email string, token string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	if !user.MFAEnabled {
		valid := totp.Validate(token, user.MFASecret)
		if !valid {
			logging.FromContext(ctx).Warn("invalid MFA token", "email", email)
			return errors.New("Invalid MFA token")
		}
		enabled := true
		_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
			MFAEnabled: &enabled,
		})
		if err != nil {
			logging.FromContext(ctx).Error("failed to enable MFA", "email", email, "error", err)
			return errors.New("Failed to enable MFA")
		}
		return nil
	}
	return errors.New("MFA is already enabled")
}

func (u *UserService) GenerateMFAAssets(ctx context.Context, email string) (string, image.Image, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return "", nil, errors.New("User not found")
	}
	if user.MFAEnabled {
		logging.FromContext(ctx).Warn("MFA already enabled", "email", email)
		return "", nil, errors.New("MFA is already enabled")
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "templateApp",
		AccountName: email,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to generate MFA key", "email", email, "error", err)
		return "", nil, errors.New("Failed to generate MFA key")
	}
	secret := key.Secret()
	_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
		MFASecret: &secret,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to enable MFA", "email", email, "error", err)
		return "", nil, errors.New("Failed to enable MFA")
	}
	img, _ := key.Image(200, 200)
	return secret, img, nil
}

func (u *UserService) ValidateMFAToken(ctx context.Context, email string, token string) bool {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return false
	}
	if !user.MFAEnabled {
		logging.FromContext(ctx).Warn("MFA not enabled", "email", email)
		return false
	}
	return totp.Validate(token, user.MFASecret)
}

func (u *UserService) DisableMFA(ctx context.Context, email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	if user.MFAEnabled {
		enabled := false
		empty := ""
		_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
			MFAEnabled: &enabled,
			MFASecret:  &empty,
		})
		if err != nil {
			logging.FromContext(ctx).Error("failed to disable MFA", "email", email, "error", err)
			return errors.New("Failed to disable MFA")
		}
		return nil
	}
	return errors.New("MFA is already disabled")
}

func (u *UserService) UpdatePassword(ctx context.Context, email string, newPassword string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := u.dbConn.UserGetByEmail(ctx, email)
	if err != nil {
		logging.FromContext(ctx).Warn("failed to locate user", "email", email)
		return errors.New("User not found")
	}
	if !password_handling.IsPasswordValid(newPassword, u.minPasswordLen, u.complexPasswords) {
		return errors.New("Invalid password")
	}
	hash := password_handling.HashAndSaltPassword(newPassword)
	_, err = u.dbConn.UserUpdate(ctx, user.ID, store.UpdateUserOptions{
		PasswordHash: &hash,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to update password", "email", email, "error", err)
		return errors.New("Failed to update password")
	}
	return nil
}

func (u *UserService) QueryUsers(ctx context.Context, options store.UserQueryOptions) ([]models.User, error) {
	users, err := u.dbConn.UserQuery(ctx, options, 100, 0)
	if err != nil {
		logging.FromContext(ctx).Error("failed to query users", "error", err)
		return nil, errors.New("Failed to query users")
	}
	listing := make([]models.User, 0)
	for _, user := range users {
		listing = append(listing, *user)
	}
	return listing, nil
}

func (u *UserService) ListUsers(ctx context.Context, pageNumber int, pageSize int) ([]models.User, error) {
	users, err := u.dbConn.UserList(ctx, pageSize, pageNumber*pageSize)
	if err != nil {
		logging.FromContext(ctx).Error("failed to list users", "error", err)
		return nil, errors.New("Failed to list users")
	}
	listing := make([]models.User, 0)
	for _, user := range users {
		listing = append(listing, *user)
	}
	return listing, nil
}

func (u *UserService) CreateUser(ctx context.Context, email string, role string, initialPassword string) (*models.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	// Make sure the email isn't already taken
	_, err := u.dbConn.UserGetByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("Email already in use")
	}
	if !password_handling.IsPasswordValid(initialPassword, u.minPasswordLen, u.complexPasswords) {
		return nil, errors.New("Invalid password")
	}
	hash := password_handling.HashAndSaltPassword(initialPassword)
	user, err := u.dbConn.UserCreate(ctx, store.CreateUserOptions{
		Email:        email,
		PasswordHash: hash,
		Role:         role,
		Invited:      true,
	})
	if err != nil {
		logging.FromContext(ctx).Error("failed to create user", "email", email, "error", err)
		return nil, errors.New("Failed to create user")
	}
	return user, nil
}

func (u *UserService) UpdateUser(ctx context.Context, id uuid.UUID, options store.UpdateUserOptions) error {
	_, err := u.dbConn.UserUpdate(ctx, id, options)
	if err != nil {
		logging.FromContext(ctx).Error("failed to update user", "id", id, "error", err)
		return errors.New("Failed to update user")
	}
	return nil
}
