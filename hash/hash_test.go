package hash

import (
	"bytes"
	"testing"
)

func TestExtractSalt(t *testing.T) {
	expectedSalt := []byte("salt")
	expectedSaltLength := 4

	hash, err := Hash([]byte("1234"), expectedSalt)
	if err != nil {
		t.Errorf("Hash returned error %v", err)
		return
	}

	actualSalt, err := ExtractSalt(hash)
	if err != nil {
		t.Errorf("ExtractSalt returned error %v", err)
		return
	}

	actualSaltLength := len(actualSalt)
	if actualSaltLength != expectedSaltLength {
		t.Errorf("Expected salt length %d, got %d. Salt: %v", expectedSaltLength, actualSaltLength, actualSalt)
		return
	}

	if !bytes.Equal(actualSalt, expectedSalt) {
		t.Errorf("Expected salt %v, got %v", expectedSalt, actualSalt)
		return
	}
}

func TestGenerateSalt(t *testing.T) {
	saltLengths := []int{0, 16, 18, 20}

	for _, expectedSaltLength := range saltLengths {
		actualSalt, err := GenerateSalt(expectedSaltLength)

		// Default salt length is 16
		if expectedSaltLength == 0 {
			expectedSaltLength = 16
		}

		if err != nil {
			t.Errorf("GenerateSalt returned error %v", err)
			return
		}

		actualSaltLength := len(actualSalt)
		if actualSaltLength != expectedSaltLength {
			t.Errorf("Expected salt length %d, got %d. Salt: %v", expectedSaltLength, actualSaltLength, actualSalt)
			return
		}
	}
}

func TestHash(t *testing.T) {
	expectedSalt := []byte("salt")

	hash, err := Hash([]byte("1234"), expectedSalt)
	if err != nil {
		t.Errorf("Hash returned error %v", err)
		return
	}

	actualSalt, err := ExtractSalt(hash)
	if err != nil {
		t.Errorf("ExtractSalt returned error %v", err)
		return
	}

	if !bytes.Equal(actualSalt, expectedSalt) {
		t.Errorf("Expected salt %v, got %v", expectedSalt, actualSalt)
	}
}

func TestHashValidate(t *testing.T) {
	passwordA := []byte("1234")
	salt := []byte("salt")
	hashA, err := Hash(passwordA, salt)
	if err != nil {
		t.Errorf("Hash returned error %v", err)
		return
	}

	passwordB := []byte("4321")
	hashB, err := Hash(passwordB, salt)
	if err != nil {
		t.Errorf("Hash returned error %v", err)
		return
	}

	valid, err := HashValidate(passwordA, hashA)
	if err != nil {
		t.Errorf("HashValidate returned error %v", err)
		return
	}
	if !valid {
		t.Error("Expected true, got false")
		return
	}

	valid, err = HashValidate(passwordA, hashB)
	if err != nil {
		t.Errorf("HashValidate returned error %v", err)
		return
	}
	if valid {
		t.Error("Expected false, got true")
		return
	}

	valid, err = HashValidate(passwordB, hashA)
	if err != nil {
		t.Errorf("HashValidate returned error %v", err)
		return
	}
	if valid {
		t.Error("Expected false, got true")
		return
	}
}
