package storage

import (
	"errors"
)

type diskStorage struct {
}

func NewDiskStorage() Storage {
	return &diskStorage{}
}

func (d *diskStorage) SaveData(data string) error {
	return errors.New("Not implemented")
}

func (d *diskStorage) GetData() (string, error) {
	return "", errors.New("Not implemented")
}
