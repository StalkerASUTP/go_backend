package factory

import (
	"fmt"
	"go-backend/internal/calculator"
	"go-backend/internal/http"
	"go-backend/internal/interfaces"
	"go-backend/internal/logger"
	"go-backend/internal/reader"
	"go-backend/internal/writer"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateDataReader(mode, filename string) (interfaces.DataReader, error) {
	switch mode {
	case "json":
		return reader.NewFileReader(filename), nil
	case "stdin":
		return reader.NewStdinReader(), nil
	default:
		return nil, fmt.Errorf("unsupported reader mode: %s", mode)
	}
}

func (f *Factory) CreateDataWriter(filename string) (interfaces.DataWriter, error) {
	return writer.NewFileWriter(filename)
}

func (f *Factory) CreateLogger(filename string) (interfaces.Logger, error) {
	return logger.NewStructuredLogger(filename)
}

func (f *Factory) CreateCalculator() interfaces.Calculator {
	return calculator.NewCalculator()
}

func (f *Factory) CreateHTTPClient() interfaces.HTTPClient {
	return http.NewClient()
}