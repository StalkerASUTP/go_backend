package service

import (
	"fmt"
	"go-backend/internal/interfaces"
	"net/http"
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
		return fmt.Errorf("URL is empty")
	}
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request to %s: %w", url, err)
	}
	defer resp.Body.Close()
	host := resp.Request.Host
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)
	if statusCode == http.StatusOK {
		p.logger.Info("HTTP request completed successfully",
			"host", host,
			"status_code", statusCode,
			"url", url)

		fmt.Printf("✅ SUCCESS: Host: %s, Status: %d OK\n", host, statusCode)
		return nil
	}

	fmt.Printf("❌ FAILED: Host: %s, Status: %d %s\n", host, statusCode, statusText)

	return fmt.Errorf("HTTP request failed: status %d (%s) for URL: %s",
		statusCode, statusText, url)

}
