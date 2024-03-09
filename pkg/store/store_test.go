package store

import (
	"context"
	"go-monolith-template/pkg/password_handling"
)

func createTestDatabase() *Storage {
	s, err := ConnectSQLITE3(SqliteAdapterOptions{
		InMemory: true,
	}, true)
	if err != nil {
		panic(err)
	}
	return NewStorage(s)
}

func createSeededTestDatabase() *Storage {
	s := createTestDatabase()
	// seed the database
	for _, u := range seedUsers {
		_, err := s.UserCreate(context.TODO(), u)
		if err != nil {
			panic(err)
		}

	}
	return s
}

var seedUsers = []CreateUserOptions{
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
