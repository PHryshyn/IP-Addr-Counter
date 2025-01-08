package processor

import (
	"math/bits"
)

type Storage struct {
	array []uint64
}

func NewStorage(size uint) *Storage {
	return &Storage{
		array: make([]uint64, size),
	}
}

func ProcessIndex(ipIndex uint64, storage *Storage) {
	var index = ipIndex / 64
	var position = ipIndex % 64

	storage.array[index] = storage.array[index] | (1 << position)
}

func CalculateNumberOfUniqueIds(storage *Storage) uint64 {
	var result uint64 = 0

	for _, value := range storage.array {
		result += uint64(bits.OnesCount64(value))
	}

	return result
}
