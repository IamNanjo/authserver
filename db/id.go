package db

import (
	"crypto/rand"
	"encoding/base64"
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
