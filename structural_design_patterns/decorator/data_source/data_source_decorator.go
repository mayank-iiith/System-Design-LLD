package datasource

// DataSourceDecorator is the base decorator embedding another DataSource.
type DataSourceDecorator struct {
	wrappee DataSource
}

func (d *DataSourceDecorator) WriteData(data string) error {
	return d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() (string, error) {
	return d.wrappee.ReadData()
}
