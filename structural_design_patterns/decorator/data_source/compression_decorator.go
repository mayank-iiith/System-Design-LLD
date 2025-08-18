package datasource

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

type CompressionDecorator struct {
	*DataSourceDecorator
	compressionLevel int
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{
		DataSourceDecorator: &DataSourceDecorator{
			wrappee: source,
		},
		compressionLevel: flate.DefaultCompression,
	}
}

func (c *CompressionDecorator) WriteData(data string) error {
	compressed, err := c.compress(data)
	if err != nil {
		return err
	}

	return c.wrappee.WriteData(compressed)
}

func (c *CompressionDecorator) ReadData() (string, error) {
	compressed, err := c.wrappee.ReadData()
	if err != nil {
		return "", err
	}

	return c.decompress(compressed)
}

func (c *CompressionDecorator) compress(data string) (string, error) {
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, c.compressionLevel)
	if err != nil {
		return "", err
	}

	_, err = writer.Write([]byte(data))
	if err != nil {
		return "", err
	}

	writer.Close()

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (c *CompressionDecorator) decompress(data string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	reader := flate.NewReader(bytes.NewReader(raw))
	defer reader.Close()

	out, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
