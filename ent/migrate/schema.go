// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "mfa_secret", Type: field.TypeString, Default: ""},
		{Name: "mfa_enabled", Type: field.TypeBool, Default: false},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "invited", Type: field.TypeBool, Default: true},
		{Name: "locked", Type: field.TypeBool, Default: false},
		{Name: "locked_at", Type: field.TypeTime, Nullable: true},
		{Name: "api_key", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}