// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-monolith-template/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash,omitempty"`
	// MfaSecret holds the value of the "mfa_secret" field.
	MfaSecret string `json:"mfa_secret,omitempty"`
	// MfaEnabled holds the value of the "mfa_enabled" field.
	MfaEnabled bool `json:"mfa_enabled,omitempty"`
	// LastLogin holds the value of the "last_login" field.
	LastLogin time.Time `json:"last_login,omitempty"`
	// Invited holds the value of the "invited" field.
	Invited bool `json:"invited,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked bool `json:"locked,omitempty"`
	// LockedAt holds the value of the "locked_at" field.
	LockedAt time.Time `json:"locked_at,omitempty"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"api_key,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// PasswordResetTokenExpiration holds the value of the "password_reset_token_expiration" field.
	PasswordResetTokenExpiration *time.Time `json:"password_reset_token_expiration,omitempty"`
	// PasswordResetToken holds the value of the "password_reset_token" field.
	PasswordResetToken string `json:"password_reset_token,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserSession holds the value of the user_session edge.
	UserSession []*Session `json:"user_session,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserSessionOrErr returns the UserSession value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserSessionOrErr() ([]*Session, error) {
	if e.loadedTypes[0] {
		return e.UserSession, nil
	}
	return nil, &NotLoadedError{edge: "user_session"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldMfaEnabled, user.FieldInvited, user.FieldLocked:
			values[i] = new(sql.NullBool)
		case user.FieldEmail, user.FieldPasswordHash, user.FieldMfaSecret, user.FieldAPIKey, user.FieldRole, user.FieldPasswordResetToken:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldLastLogin, user.FieldLockedAt, user.FieldPasswordResetTokenExpiration:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldMfaSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mfa_secret", values[i])
			} else if value.Valid {
				u.MfaSecret = value.String
			}
		case user.FieldMfaEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field mfa_enabled", values[i])
			} else if value.Valid {
				u.MfaEnabled = value.Bool
			}
		case user.FieldLastLogin:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login", values[i])
			} else if value.Valid {
				u.LastLogin = value.Time
			}
		case user.FieldInvited:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invited", values[i])
			} else if value.Valid {
				u.Invited = value.Bool
			}
		case user.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				u.Locked = value.Bool
			}
		case user.FieldLockedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field locked_at", values[i])
			} else if value.Valid {
				u.LockedAt = value.Time
			}
		case user.FieldAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value.Valid {
				u.APIKey = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = value.String
			}
		case user.FieldPasswordResetTokenExpiration:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field password_reset_token_expiration", values[i])
			} else if value.Valid {
				u.PasswordResetTokenExpiration = new(time.Time)
				*u.PasswordResetTokenExpiration = value.Time
			}
		case user.FieldPasswordResetToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_reset_token", values[i])
			} else if value.Valid {
				u.PasswordResetToken = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserSession queries the "user_session" edge of the User entity.
func (u *User) QueryUserSession() *SessionQuery {
	return NewUserClient(u.config).QueryUserSession(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password_hash=")
	builder.WriteString(u.PasswordHash)
	builder.WriteString(", ")
	builder.WriteString("mfa_secret=")
	builder.WriteString(u.MfaSecret)
	builder.WriteString(", ")
	builder.WriteString("mfa_enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.MfaEnabled))
	builder.WriteString(", ")
	builder.WriteString("last_login=")
	builder.WriteString(u.LastLogin.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("invited=")
	builder.WriteString(fmt.Sprintf("%v", u.Invited))
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", u.Locked))
	builder.WriteString(", ")
	builder.WriteString("locked_at=")
	builder.WriteString(u.LockedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("api_key=")
	builder.WriteString(u.APIKey)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(u.Role)
	builder.WriteString(", ")
	if v := u.PasswordResetTokenExpiration; v != nil {
		builder.WriteString("password_reset_token_expiration=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("password_reset_token=")
	builder.WriteString(u.PasswordResetToken)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
