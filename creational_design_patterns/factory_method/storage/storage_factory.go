// Factory method creational design pattern allows creating objects without having to specify the exact type of the object that will be created.
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
