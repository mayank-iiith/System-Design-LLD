package datasource

import (
	"encoding/base64"
)

type EncryptionDecorator struct {
	*DataSourceDecorator
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		&DataSourceDecorator{
			wrappee: source,
		},
	}
}

func (e *EncryptionDecorator) WriteData(data string) error {
	encoded := e.encode(data)

	return e.wrappee.WriteData(encoded)
}

func (e *EncryptionDecorator) ReadData() (string, error) {
	encrypted, err := e.wrappee.ReadData()
	if err != nil {
		return "", err
	}

	return e.decode(encrypted)
}

func (e *EncryptionDecorator) encode(data string) string {
	b := []byte(data)

	for i := range b {
		b[i] += 1
	}

	return base64.StdEncoding.EncodeToString(b)
}

func (e *EncryptionDecorator) decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] -= 1
	}

	return string(b), nil
}
