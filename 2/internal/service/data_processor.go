package service

import (
	"fmt"
	"go-backend/internal/interfaces"
)

type DataProcessor struct {
	reader     interfaces.DataReader
	logger     interfaces.Logger
	calculator interfaces.Calculator
	httpClient interfaces.HTTPClient
}

func NewDataProcessor(
	reader interfaces.DataReader,
	logger interfaces.Logger,
	calculator interfaces.Calculator,
	httpClient interfaces.HTTPClient,
) *DataProcessor {
	return &DataProcessor{
		reader:     reader,
		logger:     logger,
		calculator: calculator,
		httpClient: httpClient,
	}
}

func (p *DataProcessor) ProcessData() error {
	numbers, err := p.reader.ReadData()
	if err != nil {
		p.logger.Error("failed to read data", "error", err)
		return fmt.Errorf("failed to read data: %w", err)
	}
	p.logger.Info("data successfully read", "count", len(numbers))
	sum := p.calculator.Sum(numbers)
	p.logger.Info("sum calculated", "sum", sum)
	result := fmt.Sprintf("Total sum of numbers is: %d\n", sum)
	fmt.Print(result)
	return nil
}

func (p *DataProcessor) ProcessHTTPRequest(url string) error {
	if url == "" {
		err := fmt.Errorf("URL is empty")
		p.logger.Error("failed to make HTTP request", "error", err)
		return err
	}
	resp, err := p.httpClient.Get(url)
	if err != nil {
		p.logger.Error("failed to make HTTP request", "url", url, "error", err)
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()
	p.logger.Info("HTTP request completed", "host", resp.Request.Host, "status_code", resp.StatusCode)
	result := fmt.Sprintf("The response is getted. Host: %s, status code: %d\n", resp.Request.Host, resp.StatusCode)
	fmt.Print(result)
	return nil
}