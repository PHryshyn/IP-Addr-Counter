package calculator

import (
	"testing"
)

func TestCalculateIndexReturnZero(t *testing.T) {
	var ip = "0.0.0.0"

	var result = CalculateIndex(ip)

	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestCalculateIndexReturnMax(t *testing.T) {
	var ip = "255.255.255.255"

	var result = CalculateIndex(ip)

	if result != 4294967295 {
		t.Errorf("Expected 4294967295, but got %d", result)
	}
}

func TestCalculateIndexReturnCorrect(t *testing.T) {
	var ip = "103.50.2.1"
	var result = CalculateIndex(ip)

	if result != (256*256*256*103 + 256*256*50 + 256*2 + 1) {
		t.Errorf("Got incorrect result:  %d", result)
	}
}
