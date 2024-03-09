package testing_helpers

import (
	"context"
	"go-monolith-template/pkg/password_handling"
	"go-monolith-template/pkg/store"
)

func CreateTestDatabase() *store.Storage {
	s, err := store.ConnectSQLITE3(store.SqliteAdapterOptions{
		InMemory: true,
	}, true)
	if err != nil {
		panic(err)
	}
	return store.NewStorage(s)
}

func CreateSeededTestDatabase() *store.Storage {
	s := CreateTestDatabase()
	// seed the database
	for _, u := range seedUsers {
		_, err := s.UserCreate(context.TODO(), u)
		if err != nil {
			panic(err)
		}

	}
	return s
}

var seedUsers = []store.CreateUserOptions{
	{
		Email:        "admin1@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "admin",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
	{
		Email:        "admin2@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "admin",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
	{
		Email:        "admin3@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "admin",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
	{
		Email:        "user1@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "user",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
	{
		Email:        "user2@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "user",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
	{
		Email:        "user3@example.com",
		PasswordHash: password_handling.HashAndSaltPassword("password"),
		Role:         "user",
		MFASecret:    "",
		MFARequired:  false,
		APIKey:       password_handling.GenerateRandomPassword(32),
		Invited:      false,
	},
}
