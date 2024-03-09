package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Metadata is a mixin that adds metadata fields to a schema
type Metadata struct {
	ent.Mixin
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return nil
}

// Indexes of the Metadata.
func (m Metadata) Indexes() []ent.Index {
	return nil
}

// Hooks of the Metadata.
func (m Metadata) Hooks() []ent.Hook {
	return nil
}

// Interceptors of the Metadata.
func (m Metadata) Interceptors() []ent.Interceptor {
	return nil
}

// Policy of the Metadata.
func (m Metadata) Policy() ent.Policy {
	return nil
}

// Annotations of the Metadata.
func (m Metadata) Annotations() []schema.Annotation {
	return nil
}
