package counter

import (
	"path/filepath"
	"testing"
)

func TestReadFileAncCalculateUniqueIds(t *testing.T) {
	filePath := filepath.Join("test_data", "example.zip")

	var result, err = ReadFileAncCalculateUniqueIds(filePath)

	if err != nil {
		t.Errorf("Error during processing: %v", err)
	}

	if result != 100 {
		t.Errorf("Expected not zero, but got %d", result)
	}
}
