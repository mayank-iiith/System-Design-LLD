package datasource

import (
	"os"
)

type FileDataSource struct {
	fileName string
}

func NewFileDataSource(fileName string) *FileDataSource {
	return &FileDataSource{
		fileName: fileName,
	}
}

func (f *FileDataSource) WriteData(data string) error {
	os.WriteFile(f.fileName, []byte(data), 0644)
	return nil
}

func (f *FileDataSource) ReadData() (string, error) {
	bytes, err := os.ReadFile(f.fileName)
	return string(bytes), err
}
