package hash

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

func ExtractSalt(hash string) ([]byte, error) {
	parts := strings.Split(hash, "$")

	if len(parts) < 5 {
		return nil, fmt.Errorf("Invalid argon2 hash. Too few parts")
	}

	decodedSalt, err := base64.StdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, err
	}

	return decodedSalt, nil
}

func ExtractHashKey(hash string) ([]byte, error) {
	parts := strings.Split(hash, "$")

	if len(parts) < 5 {
		return nil, fmt.Errorf("Invalid argon2 hash. Too few parts")
	}

	decodedHashKey, err := base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, err
	}

	return decodedHashKey, nil
}

func GenerateSalt(saltLength int) ([]byte, error) {
	// Set default salt length
	if saltLength == 0 {
		saltLength = 16
	}

	salt := make([]byte, saltLength)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Automatically generates a salt if nil is passed
func Hash(password []byte, salt []byte) (string, error) {
	saltLength := 16
	actualSaltLength := len(salt)

	if salt != nil && len(salt) != actualSaltLength {
		return "", fmt.Errorf("Expected salt length %d, got %d", saltLength, actualSaltLength)
	} else if salt == nil {
		saltGenerated, error := GenerateSalt(saltLength)
		if error != nil {
			return "", error
		}
		salt = saltGenerated
	}

	if password == nil || len(password) == 0 {
		return "", fmt.Errorf("Password cannot be empty")
	}

	if salt == nil || len(salt) == 0 {
		return "", fmt.Errorf("Salt cannot be empty")
	}

	var timeCost uint32 = 1
	var memoryCost uint32 = 64 * 1024
	var paralellism uint8 = 1
	var keyLength uint32 = 33

	hashKey := argon2.IDKey(password, salt, timeCost, memoryCost, paralellism, keyLength)

	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedKey := base64.StdEncoding.EncodeToString(hashKey)

	hash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", memoryCost, timeCost, paralellism, encodedSalt, encodedKey)

	return hash, nil
}

func HashValidate(password []byte, hash string) (bool, error) {
	salt, err := ExtractSalt(hash)
	if err != nil {
		return false, err
	}

	hashedPassword, err := Hash(password, salt)
	if err != nil {
		return false, err
	}

	hashedPasswordKey, err := ExtractHashKey(hashedPassword)
	if err != nil {
		return false, err
	}

	hashKey, err := ExtractHashKey(hash)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(hashedPasswordKey, hashKey) {
		return false, nil
	}

	return true, nil
}
