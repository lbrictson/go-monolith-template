package store

import (
	"context"
	"github.com/google/uuid"
	"go-monolith-template/ent"
	"go-monolith-template/ent/user"
	"go-monolith-template/pkg/models"
	"time"
)

func convertEntUserModelToDomainUserModel(entUser *ent.User) *models.User {
	u := models.User{
		ID:           entUser.ID,
		Email:        entUser.Email,
		PasswordHash: entUser.PasswordHash,
		Locked:       entUser.Locked,
		APIKey:       entUser.APIKey,
		MFAEnabled:   entUser.MfaEnabled,
		MFASecret:    entUser.MfaSecret,
		Invited:      entUser.Invited,
		CreatedAt:    entUser.CreatedAt,
		UpdatedAt:    entUser.UpdatedAt,
		Role:         entUser.Role,
	}
	if u.Locked {
		u.LockedAt = &entUser.LockedAt
	} else {
		u.LockedAt = nil
	}
	if !u.Invited {
		u.LastLogin = &entUser.LastLogin
	} else {
		u.LastLogin = nil
	}
	return &u
}

type CreateUserOptions struct {
	Email        string
	PasswordHash string
	MFARequired  bool
	MFASecret    string
	APIKey       string
	Invited      bool
	Role         string
}

func (s *Storage) UserCreate(ctx context.Context, opts CreateUserOptions) (*models.User, error) {
	entUser, err := s.conn.User.Create().
		SetEmail(opts.Email).
		SetPasswordHash(opts.PasswordHash).
		SetMfaEnabled(opts.MFARequired).
		SetMfaSecret(opts.MFASecret).
		SetAPIKey(opts.APIKey).
		SetInvited(opts.Invited).
		SetRole(opts.Role).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntUserModelToDomainUserModel(entUser), nil
}

func (s *Storage) UserGetByEmail(ctx context.Context, email string) (*models.User, error) {
	entUser, err := s.conn.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntUserModelToDomainUserModel(entUser), nil
}

func (s *Storage) UserGetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	entUser, err := s.conn.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntUserModelToDomainUserModel(entUser), nil
}

type UpdateUserOptions struct {
	PasswordHash   *string
	Locked         *bool
	APIKey         *string
	MFAEnabled     *bool
	MFASecret      *string
	LastLogin      *time.Time
	LockedAt       *time.Time
	Invited        *bool
	ClearLastLogin bool
	ClearLockedAt  bool
	Role           *string
}

func (s *Storage) UserUpdate(ctx context.Context, id uuid.UUID, opts UpdateUserOptions) (*models.User, error) {
	update := s.conn.User.UpdateOneID(id)
	if opts.PasswordHash != nil {
		update = update.SetPasswordHash(*opts.PasswordHash)
	}
	if opts.Locked != nil {
		update = update.SetLocked(*opts.Locked)
	}
	if opts.APIKey != nil {
		update = update.SetAPIKey(*opts.APIKey)
	}
	if opts.MFAEnabled != nil {
		update = update.SetMfaEnabled(*opts.MFAEnabled)
	}
	if opts.MFASecret != nil {
		update = update.SetMfaSecret(*opts.MFASecret)
	}
	if opts.LastLogin != nil {
		update = update.SetLastLogin(*opts.LastLogin)
	}
	if opts.LockedAt != nil {
		update = update.SetLockedAt(*opts.LockedAt)
	}
	if opts.Invited != nil {
		update = update.SetInvited(*opts.Invited)
	}
	if opts.ClearLastLogin {
		update = update.ClearLastLogin()
	}
	if opts.ClearLockedAt {
		update = update.ClearLockedAt()
	}
	if opts.Role != nil {
		update = update.SetRole(*opts.Role)
	}
	entUser, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntUserModelToDomainUserModel(entUser), nil
}

func (s *Storage) UserDelete(ctx context.Context, id uuid.UUID) error {
	return s.conn.User.DeleteOneID(id).Exec(ctx)
}

func (s *Storage) UserList(ctx context.Context, limit int, offset int) ([]*models.User, error) {
	entUsers, err := s.conn.User.Query().Limit(limit).Offset(offset).Order(ent.Asc(user.FieldEmail)).All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(entUsers))
	for _, u := range entUsers {
		users = append(users, convertEntUserModelToDomainUserModel(u))
	}
	return users, nil
}

type UserQueryOptions struct {
	EmailLike       *string
	Locked          *bool
	Invited         *bool
	LastLoginBefore *time.Time
	LastLoginAfter  *time.Time
	LockedBefore    *time.Time
	LockedAfter     *time.Time
	RoleIs          *string
}

func (s *Storage) UserQuery(ctx context.Context, opts UserQueryOptions, limit int, offset int) ([]*models.User, error) {
	query := s.conn.User.Query()
	if opts.EmailLike != nil {
		query = query.Where(user.EmailContains(*opts.EmailLike))
	}
	if opts.Locked != nil {
		query = query.Where(user.Locked(*opts.Locked))
	}
	if opts.Invited != nil {
		query = query.Where(user.Invited(*opts.Invited))
	}
	if opts.LastLoginBefore != nil {
		query = query.Where(user.LastLoginLT(*opts.LastLoginBefore))
	}
	if opts.LastLoginAfter != nil {
		query = query.Where(user.LastLoginGT(*opts.LastLoginAfter))
	}
	if opts.LockedBefore != nil {
		query = query.Where(user.LockedAtLT(*opts.LockedBefore))
	}
	if opts.LockedAfter != nil {
		query = query.Where(user.LockedAtGT(*opts.LockedAfter))
	}
	if opts.RoleIs != nil {
		query = query.Where(user.Role(*opts.RoleIs))
	}
	query = query.Limit(limit)
	query = query.Offset(offset)
	entUsers, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(entUsers))
	for _, u := range entUsers {
		users = append(users, convertEntUserModelToDomainUserModel(u))
	}
	return users, nil
}
