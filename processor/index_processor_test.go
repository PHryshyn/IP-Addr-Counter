package processor

import (
	"testing"
)

func TestProcessIndexWhenIpIndexIsZero(t *testing.T) {
	var storage = NewStorage(5)
	var IpIndex uint64 = 0
	var expected uint64 = 0b00001
	var expectedIndex = IpIndex / 64

	ProcessIndex(IpIndex, storage)

	if storage.array[expectedIndex] != expected {
		t.Errorf("Expected 1, but got %d", storage.array[0])
	}
}

func TestProcessIndexWhenIpIndexIs63(t *testing.T) {
	var storage = NewStorage(5)
	var IpIndex uint64 = 63
	var expected uint64 = 1 << (IpIndex % 64)
	var expectedIndex = IpIndex / 64

	ProcessIndex(IpIndex, storage)

	if storage.array[expectedIndex] != expected {
		t.Errorf("Expected uint64 %d, but got %d", expected, storage.array[expectedIndex])
	}
}

func TestProcessIndexWhenIpIndexIs64(t *testing.T) {
	var storage = NewStorage(5)
	var IpIndex uint64 = 64
	var expected uint64 = 1 << IpIndex % 64
	var expectedIndex = IpIndex / 64

	ProcessIndex(IpIndex, storage)

	if storage.array[expectedIndex] != expected && storage.array[0] != 0 {
		t.Errorf("Expected 1, but got %d", storage.array[expectedIndex])
	}
}

func TestProcessIndexWhenIpIndexIs150(t *testing.T) {
	var storage = NewStorage(5)
	var IpIndex uint64 = 150
	var expected uint64 = 1 << IpIndex % 64
	var expectedIndex = IpIndex / 64

	ProcessIndex(IpIndex, storage)

	if storage.array[expectedIndex] != expected && storage.array[0] != 0 && storage.array[1] != 0 {
		t.Errorf("Expected 1, but got %d", storage.array[expectedIndex])
	}
}

func TestProcessIndexWhenIpIndexIs150RunTwice(t *testing.T) {
	var storage = NewStorage(5)
	var IpIndex uint64 = 150
	var expected uint64 = 1 << IpIndex % 64
	var expectedIndex = IpIndex / 64

	ProcessIndex(IpIndex, storage)
	ProcessIndex(IpIndex, storage)

	if storage.array[expectedIndex] != expected && storage.array[0] != 0 && storage.array[1] != 0 {
		t.Errorf("Expected 1, but got %d", storage.array[expectedIndex])
	}
}

func TestCalculateNumberOfUniqueIdsReturnZero(t *testing.T) {
	var storage = NewStorage(5)
	var result = CalculateNumberOfUniqueIds(storage)

	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestCalculateNumberOfUniqueIds(t *testing.T) {
	var storage = NewStorage(5)

	ProcessIndex(10, storage)
	ProcessIndex(11, storage)
	ProcessIndex(10, storage)
	ProcessIndex(30, storage)
	ProcessIndex(20, storage)
	ProcessIndex(1, storage)
	ProcessIndex(3, storage)
	ProcessIndex(15, storage)
	ProcessIndex(4, storage)
	ProcessIndex(7, storage)

	var result = CalculateNumberOfUniqueIds(storage)

	if result != 9 {
		t.Errorf("Expected 9, but got %d", result)
	}
}
