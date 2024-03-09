package core

import (
	"context"
	"go-monolith-template/pkg/store"
	"go-monolith-template/pkg/testing_helpers"
	"testing"
)

func TestUserService_UserLockedOut(t *testing.T) {
	db := testing_helpers.CreateSeededTestDatabase()
	userService := NewUserService(UserServiceOptions{
		StorageLayer:             db,
		LockoutThreshold:         3,
		ComplexPasswordsRequired: true,
		MinPasswordLength:        8,
		MFAMandatory:             false,
	})
	// Try to login with wrong password 3 times
	_, _, _ = userService.AuthenticateUser(context.TODO(), "admin1@example.com", "wrongpassword")
	_, _, _ = userService.AuthenticateUser(context.TODO(), "admin1@example.com", "wrongpassword")
	_, _, _ = userService.AuthenticateUser(context.TODO(), "admin1@example.com", "wrongpassword")
	// Now the user should be locked out
	lockedUser, err := db.UserGetByEmail(context.TODO(), "admin1@example.com")
	if err != nil {
		t.Error(err)
	}
	if !lockedUser.Locked {
		t.Error("User should be locked out")
	}
	// Try to login with the correct password, it should fail because the account is locked
	_, _, err = userService.AuthenticateUser(context.TODO(), "admin1@example.com", "password")
	if err == nil {
		t.Error("User should be locked out")
	}
}

func TestUserService_AuthenticateUser(t *testing.T) {
	type fields struct {
		dbConn           *store.Storage
		lockoutTracker   map[string]int
		lockoutThreshold int
		minPasswordLen   int
		complexPasswords bool
		mfaMandatory     bool
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   bool
		wantErr bool
	}{
		{
			name: "Test AuthenticateUser:  Happy Path - no MFA",
			fields: fields{
				dbConn:           testing_helpers.CreateSeededTestDatabase(),
				lockoutTracker:   make(map[string]int),
				lockoutThreshold: 3,
				minPasswordLen:   8,
				complexPasswords: true,
				mfaMandatory:     false,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "admin1@example.com",
				password: "password",
			},
			want:    true,
			want1:   false,
			wantErr: false,
		},
		{
			name: "Test AuthenticateUser:  Sad Path wrong password",
			fields: fields{
				dbConn:           testing_helpers.CreateSeededTestDatabase(),
				lockoutTracker:   make(map[string]int),
				lockoutThreshold: 3,
				minPasswordLen:   8,
				complexPasswords: true,
				mfaMandatory:     true,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "admin1@example.com",
				password: "wrongpassword",
			},
			want:    false,
			want1:   false,
			wantErr: true,
		},
		{
			name: "Test AuthenticateUser:  Sad Path - User does not exist",
			fields: fields{
				dbConn:           testing_helpers.CreateSeededTestDatabase(),
				lockoutTracker:   make(map[string]int),
				lockoutThreshold: 3,
				minPasswordLen:   8,
				complexPasswords: true,
				mfaMandatory:     true,
			},
			args: args{
				ctx:      context.TODO(),
				email:    "notreal@example.com",
				password: "password",
			},
			want:    false,
			want1:   false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				dbConn:           tt.fields.dbConn,
				lockoutTracker:   tt.fields.lockoutTracker,
				lockoutThreshold: tt.fields.lockoutThreshold,
				minPasswordLen:   tt.fields.minPasswordLen,
				complexPasswords: tt.fields.complexPasswords,
				mfaMandatory:     tt.fields.mfaMandatory,
			}
			got, got1, err := u.AuthenticateUser(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthenticateUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AuthenticateUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
