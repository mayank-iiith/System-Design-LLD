package datasource

// DataSource defines the interface for reading/writing data.
type DataSource interface {
	WriteData(data string) error
	ReadData() (string, error)
}
