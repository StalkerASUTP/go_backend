package factory

import (
	"go-backend/internal/calculator"
	"go-backend/internal/http"
	"go-backend/internal/interfaces"
	"go-backend/internal/logger"
	"go-backend/internal/reader"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateDataReader(filename string) (interfaces.DataReader) {
	return reader.NewFileReader(filename)

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
