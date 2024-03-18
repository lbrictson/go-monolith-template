// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMfaCompleted holds the string denoting the mfa_completed field in the database.
	FieldMfaCompleted = "mfa_completed"
	// EdgeUserSession holds the string denoting the user_session edge name in mutations.
	EdgeUserSession = "user_session"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// UserSessionTable is the table that holds the user_session relation/edge.
	UserSessionTable = "sessions"
	// UserSessionInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserSessionInverseTable = "users"
	// UserSessionColumn is the table column denoting the user_session relation/edge.
	UserSessionColumn = "user_user_session"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMfaCompleted,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sessions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_user_session",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMfaCompleted holds the default value on creation for the "mfa_completed" field.
	DefaultMfaCompleted bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByMfaCompleted orders the results by the mfa_completed field.
func ByMfaCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMfaCompleted, opts...).ToFunc()
}

// ByUserSessionField orders the results by user_session field.
func ByUserSessionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserSessionStep(), sql.OrderByField(field, opts...))
	}
}
func newUserSessionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserSessionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserSessionTable, UserSessionColumn),
	)
}