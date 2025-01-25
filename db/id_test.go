package db

import (
	"slices"
	"testing"
)

func TestGenerateId(t *testing.T) {
	lengths := []int{-5, 0, 5, 10, 15, 20, 25}

	generatedIDs := make([]string, 100)
	for _, length := range lengths {
		for i := 0; i < 100; i++ {
			id, err := GenerateId(length)
			if err != nil {
				t.Errorf("ID generation failed. %v", err)
			}

			idIsDuplicate := slices.Contains(generatedIDs, id)
			if idIsDuplicate && length > 0 {
				t.Errorf("Duplicate ID generated for length %d. %s", length, id)
			}

			generatedIDs[i] = id

			actualLength := len(id)
			if actualLength != length && (length < 1 && actualLength != 0) {
				t.Errorf("Expected ID length %d, got %d. ID: %v", length, actualLength, id)
			}
		}
	}
}
