package db

import (
	"context"
	"testing"

	"github.com/IamNanjo/authserver/hash"
)

var userId string
var testUsername = "Test"
var testEmail = "test@example.com"

func TestCreateUser(t *testing.T) {
	var err error

	hashedPassword, err := hash.Hash([]byte("1234"), nil)
	if err != nil {
		t.Fatalf("Failed to hash password 1234")
	}

	user, err := Q().CreateUser(context.Background(), CreateUserParams{
		Name:     testUsername,
		Email:    &testEmail,
		Password: &hashedPassword,
	})
	if err != nil {
		t.Fatalf("User creation failed. %v", err)
	}
	userId = user.Id

	_, err = Q().GetUserById(context.Background(), userId)
	if err != nil {
		t.Fatalf("Created user not found. %v", err)
	}
}

func TestGetUserByEmailOrUsername(t *testing.T) {
	Q().GetUserByEmailOrUsername(context.Background(), &testUsername)
}
