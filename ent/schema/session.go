package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-monolith-template/ent/mixin"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("mfa_completed").Default(false),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_session", User.Type).Unique().Ref("user_session"),
	}
}

// Mixin of the Session.
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata{},
	}
}
