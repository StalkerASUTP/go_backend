package main

import (
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/logger"
	"go-backend/internal/reader"
	"go-backend/internal/request"
	"os"
)

func sum(slice []int64) int64 {
	var sum int64
	for _, el := range slice {
		sum += el
	}
	return sum
}
func main() {
	cfg := config.LoadConfig()
	logger, err := logger.NewLogger(cfg.Logs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to inisialise logger: %s\n", err)
		return
	}

	defer logger.CloseFile()
	logger.Info("config loaded")
	logger.Info("logger inisialised")
	slice, err := reader.NewReader(cfg.Json).ReadFile()
	if err != nil {
		logger.Error("failed to read json file", "error", err)
		return
	}
	logger.Info("Numbers have been read from JSON file")
	logger.Info("Total sum of numbers is", "sum", sum(slice))
	fmt.Fprintf(os.Stdout, "Total sum of numbers is: %d\n", sum(slice))
	res, err := request.NewRequest(cfg.URL).GetResponse()
	if err != nil {
		logger.Error("failed to get response", "url", cfg.URL)
		return
	}
	logger.Info("The response is getted", "host", res.Request.Host, "status code", res.StatusCode)
	fmt.Fprintf(os.Stdout, "The response is getted. Host: %s, status code: %d\n", res.Request.Host, res.StatusCode)
}
