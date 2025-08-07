package storage

type Storage interface {
	SaveData(data string) error
	GetData() (string, error)
}
