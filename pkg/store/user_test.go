package store

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-monolith-template/ent"
	"go-monolith-template/pkg/password_handling"
	"reflect"
	"testing"
)

func TestStorage_UserCreate(t *testing.T) {
	type fields struct {
		conn *Storage
	}
	type args struct {
		ctx  context.Context
		opts CreateUserOptions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantEmail string
		wantErr   bool
	}{
		{
			name: "Test UserCreate: Happy Path",
			fields: fields{
				conn: createSeededTestDatabase(),
			},
			args: args{
				ctx: context.TODO(),
				opts: CreateUserOptions{
					Email:        "notTakenYet@example.com",
					PasswordHash: password_handling.HashAndSaltPassword("password"),
					MFARequired:  false,
					MFASecret:    "",
					APIKey:       "",
					Invited:      false,
					Role:         "Admin",
				},
			},
			wantEmail: "notTakenYet@example.com",
			wantErr:   false,
		},
		{
			name: "Test UserCreate: Email Taken",
			fields: fields{
				conn: createSeededTestDatabase(),
			},
			args: args{
				ctx: context.TODO(),
				opts: CreateUserOptions{
					Email:        "admin1@example.com",
					PasswordHash: password_handling.HashAndSaltPassword("password"),
					MFARequired:  false,
					MFASecret:    "",
					APIKey:       "",
					Invited:      false,
					Role:         "Admin",
				},
			},
			wantEmail: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.conn.UserCreate(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.Email, tt.wantEmail) {
					t.Errorf("UserCreate() got = %v, want %v", got.Email, tt.wantEmail)
				}
			}
		})
	}
}

func TestStorage_UserDelete(t *testing.T) {
	db := createSeededTestDatabase()
	happyPathUser, err := db.UserGetByEmail(context.TODO(), "admin1@example.com")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		conn *ent.Client
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test UserDelete: Happy Path",
			fields: fields{
				conn: db.conn,
			},
			args: args{
				ctx: context.TODO(),
				id:  happyPathUser.ID,
			},
			wantErr: false,
		},
		{
			name: "Test UserDelete: User Not Found",
			fields: fields{
				conn: db.conn,
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
			}
			if err := s.UserDelete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UserGetByEmail(t *testing.T) {
	type fields struct {
		conn *ent.Client
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantEmail string
		wantErr   bool
	}{
		{
			name: "Test UserGetByEmail: Happy Path",
			fields: fields{
				conn: createSeededTestDatabase().conn,
			},
			args: args{
				ctx:   context.TODO(),
				email: "admin1@example.com",
			},
			wantEmail: "admin1@example.com",
			wantErr:   false,
		},
		{
			name: "Test UserGetByEmail: User Not Found",
			fields: fields{
				conn: createSeededTestDatabase().conn,
			},
			args: args{
				ctx:   context.TODO(),
				email: fmt.Sprintf("%v@example.com", uuid.New().String()[:8]),
			},
			wantEmail: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
			}
			got, err := s.UserGetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserGetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.Email, tt.wantEmail) {
					t.Errorf("UserGetByEmail() got = %v, want %v", got, tt.wantEmail)
				}
			}
		})
	}
}

func TestStorage_UserList(t *testing.T) {
	type fields struct {
		conn *ent.Client
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test UserList: Happy Path",
			fields: fields{
				conn: createSeededTestDatabase().conn,
			},
			args: args{
				ctx:    context.TODO(),
				limit:  3,
				offset: 0,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
			}
			got, err := s.UserList(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("UserList() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}
