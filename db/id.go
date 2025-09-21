package db

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateId(length int) (string, error) {
	if length < 1 {
		return "", nil
	}

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	id := base64.RawURLEncoding.EncodeToString(randomBytes)[:length]

	return id, nil
}

func GenerateUniqueUserId(length int) (string, error) {
	return generateUniqueId(length, userIdIsUnique)
}

func GenerateUniqueSessionId(length int) (string, error) {
	return generateUniqueId(length, sessionIdIsUnique)
}

func generateUniqueId(length int, isUnique func(id string) bool) (string, error) {
	var remainingAttempts = 5
	var id = ""
	var err error

	for {
		if remainingAttempts == 0 {
			return id, errors.New("Could not create unique ID")
		}

		id, err = GenerateId(length)
		if err != nil {
			return id, err
		}

		if isUnique(id) {
			break
		}

		remainingAttempts--
	}

	return id, err
}

func userIdIsUnique(id string) bool {
	_, err := Q().GetUserById(context.Background(), id)
	return err == nil
}

func sessionIdIsUnique(id string) bool {
	_, err := Q().GetSessionById(context.Background(), id)
	return err == nil
}
