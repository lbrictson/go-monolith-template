package store

import (
	"context"
	"github.com/google/uuid"
	"go-monolith-template/ent"
	"go-monolith-template/ent/session"
	"go-monolith-template/pkg/models"
)

func convertEntSessionToModel(entSession *ent.Session) *models.SessionData {
	s := models.SessionData{
		ID:        entSession.ID,
		CreatedAt: entSession.CreatedAt,
	}
	if entSession.Edges.UserSession != nil {
		s.UserID = entSession.Edges.UserSession.ID
		s.Email = entSession.Edges.UserSession.Email
		s.Role = entSession.Edges.UserSession.Role
		s.MFAEnabled = entSession.Edges.UserSession.MfaEnabled
	}
	s.MFACompleted = entSession.MfaCompleted
	return &s
}

func (s *Storage) SessionGetByID(ctx context.Context, id uuid.UUID) (*models.SessionData, error) {
	entSession, err := s.conn.Session.Query().
		Where(session.ID(id)).
		WithUserSession().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertEntSessionToModel(entSession), nil
}

type CreateSessionOptions struct {
	UserID       uuid.UUID
	MFACompleted bool
}

func (s *Storage) SessionCreate(ctx context.Context, opts CreateSessionOptions) (*models.SessionData, error) {
	entSession, err := s.conn.Session.Create().
		SetMfaCompleted(opts.MFACompleted).
		SetUserSessionID(opts.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.SessionGetByID(ctx, entSession.ID)
}

func (s *Storage) SessionDelete(ctx context.Context, id uuid.UUID) error {
	return s.conn.Session.DeleteOneID(id).Exec(ctx)
}

type UpdateSessionOptions struct {
	MFACompleted *bool
}

func (s *Storage) SessionUpdate(ctx context.Context, id uuid.UUID, opts UpdateSessionOptions) (*models.SessionData, error) {
	builder := s.conn.Session.UpdateOneID(id)
	if opts.MFACompleted != nil {
		builder = builder.SetMfaCompleted(*opts.MFACompleted)
	}
	entSession, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.SessionGetByID(ctx, entSession.ID)
}
