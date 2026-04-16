package db

import (
	"context"
	"crypto/rand"
	"encoding/base64"
)

func GenerateId(length int) string {
	if length < 1 {
		return ""
	}

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	id := base64.RawURLEncoding.EncodeToString(randomBytes)[:length]

	return id
}

func GenerateUniqueUserId(length int) string {
	return generateUniqueId(length, userIdIsUnique)
}

func GenerateUniqueSessionId(length int) string {
	return generateUniqueId(length, sessionIdIsUnique)
}

func generateUniqueId(length int, isUnique func(id string) bool) string {
	var id = ""

	for {
		id = GenerateId(length)

		if isUnique(id) {
			break
		}
	}

	return id
}

func userIdIsUnique(id string) bool {
	_, err := Q.GetUserById(context.Background(), id)
	return err == nil
}

func sessionIdIsUnique(id string) bool {
	_, err := Q.GetSessionById(context.Background(), id)
	return err == nil
}
