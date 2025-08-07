package storage

type memoryStorage struct {
	data string
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		data: "",
	}
}

func (m *memoryStorage) SaveData(data string) error {
	m.data = data
	return nil
}

func (m *memoryStorage) GetData() (string, error) {
	return m.data, nil
}
