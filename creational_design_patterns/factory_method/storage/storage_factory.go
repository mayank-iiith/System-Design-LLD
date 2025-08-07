package storage

import (
	"fmt"
)

type StorageType int

const (
	DiskStorageType StorageType = iota + 1
	MemoryStorageType
)

func NewStorage(t StorageType) (Storage, error) {
	switch t {
	case DiskStorageType:
		return NewDiskStorage(), nil
	case MemoryStorageType:
		return NewMemoryStorage(), nil
	default:
		return nil, fmt.Errorf("invalid storage type, storage type: %d", t)
	}
}
