package db

import (
	"testing"
)

var userId string

func TestCreateUser(t *testing.T) {
	var err error

	userId, err = CreateUser("Test", "", "1234")
	if err != nil {
		t.Fatalf("User creation failed. %v", err)
	}

	_, err = GetUserById(userId)
	if err != nil {
		t.Fatalf("Created user not found. %v", err)
	}
}

func TestGetUserByEmailOrUsername(t *testing.T) {
	GetUserByEmailOrUsername("Test")
}
