package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-monolith-template/ent/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password_hash"),
		field.String("mfa_secret").Default(""),
		field.Bool("mfa_enabled").Default(false),
		field.Time("last_login").Optional(),
		field.Bool("invited").Default(true),
		field.Bool("locked").Default(false),
		field.Time("locked_at").Optional(),
		field.String("api_key"),
		field.String("role").Default("user"),
		field.Time("password_reset_token_expiration").Optional().Nillable(),
		field.String("password_reset_token").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_session", Session.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata{},
	}
}
