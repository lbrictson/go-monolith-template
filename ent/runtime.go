// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-monolith-template/ent/schema"
	"go-monolith-template/ent/session"
	"go-monolith-template/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionMixinFields0[1].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionMixinFields0[2].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sessionDescMfaCompleted is the schema descriptor for mfa_completed field.
	sessionDescMfaCompleted := sessionFields[0].Descriptor()
	// session.DefaultMfaCompleted holds the default value on creation for the mfa_completed field.
	session.DefaultMfaCompleted = sessionDescMfaCompleted.Default.(bool)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionMixinFields0[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescMfaSecret is the schema descriptor for mfa_secret field.
	userDescMfaSecret := userFields[2].Descriptor()
	// user.DefaultMfaSecret holds the default value on creation for the mfa_secret field.
	user.DefaultMfaSecret = userDescMfaSecret.Default.(string)
	// userDescMfaEnabled is the schema descriptor for mfa_enabled field.
	userDescMfaEnabled := userFields[3].Descriptor()
	// user.DefaultMfaEnabled holds the default value on creation for the mfa_enabled field.
	user.DefaultMfaEnabled = userDescMfaEnabled.Default.(bool)
	// userDescInvited is the schema descriptor for invited field.
	userDescInvited := userFields[5].Descriptor()
	// user.DefaultInvited holds the default value on creation for the invited field.
	user.DefaultInvited = userDescInvited.Default.(bool)
	// userDescLocked is the schema descriptor for locked field.
	userDescLocked := userFields[6].Descriptor()
	// user.DefaultLocked holds the default value on creation for the locked field.
	user.DefaultLocked = userDescLocked.Default.(bool)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[9].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
