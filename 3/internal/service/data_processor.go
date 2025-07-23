package service

import (
	"fmt"
	"go-backend/internal/interfaces"
	"net/http"
)

type DataProcessor struct {
	reader     interfaces.DataReader
	writer     interfaces.DataWriter
	logger     interfaces.Logger
	calculator interfaces.Calculator
	httpClient interfaces.HTTPClient
}

func NewDataProcessor(
	reader interfaces.DataReader,
	writer interfaces.DataWriter,
	logger interfaces.Logger,
	calculator interfaces.Calculator,
	httpClient interfaces.HTTPClient,
) *DataProcessor {
	return &DataProcessor{
		reader:     reader,
		writer:     writer,
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
	if err := p.writer.WriteData(result); err != nil {
		return fmt.Errorf("failed to write sum result: %w", err)
	}
	fmt.Print(result)
	return nil
}

func (p *DataProcessor) ProcessHTTPRequest(url string) error {
	if url == "" {
		return fmt.Errorf("URL is empty")
	}
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	host := resp.Request.Host
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)
	var result string
	if statusCode == http.StatusOK {
		result = fmt.Sprintf("✅ SUCCESS: Host: %s, Status: %d OK\n",
		host, statusCode)
		p.logger.Info("HTTP request completed", "host", host,
		"status_code", statusCode)
		

	} else {
		result = fmt.Sprintf("❌ FAILED: Host: %s, Status: %d %s\n",
		host, statusCode, http.StatusText(statusCode))
		p.logger.Error("HTTP request failed", "host", host,
		"status_code", statusCode, "status_text", statusText, "url", url)
	}
	if err := p.writer.WriteData(result); err != nil {
		return fmt.Errorf("failed to write HTTP result: %w", err)
	}
	fmt.Print(result)
	return nil
}
