package main

import (
	"fmt"
	"go_backend/2/internal/config"
	"go_backend/2/internal/logger"
	"go_backend/2/internal/reader"
	"go_backend/2/internal/request"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	logger, logFile, err := logger.NewLogger(cfg.Logs).InisialiseLog()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to inisialise logger: %s\n", err)
		return
	}
	defer logFile.Close()
	logger.Info("config loaded")
	fmt.Fprintf(os.Stdout, "Config loaded\n")
	logger.Info("logger inisialised")
	fmt.Fprintf(os.Stdout, "Logger inisialised\n")
	slice, err := reader.NewReader(cfg.Json).ReadFile()
	if err != nil {
		logger.Error("failed to read json file", "error", err)
		return
	}
	logger.Info("Numbers have been read from JSON file")
	fmt.Fprintf(os.Stdout, "Numbers have been read from JSON file\n")
	var sum int64
	for _, e := range slice {
		sum += e
	}
	logger.Info("Total sum of numbers is", "sum", sum)
	fmt.Fprintf(os.Stdout, "Total sum of numbers is: %d\n", sum)
	res, err := request.NewRequest(cfg.URL).GetResponse()
	if err != nil {
		logger.Error("failed to get response", "url", cfg.URL)
		return
	}
	logger.Info("The response is getted", "host", res.Request.Host, "status code", res.StatusCode)
	fmt.Fprintf(os.Stdout, "The response is getted. Host: %s, status code: %d\n", res.Request.Host, res.StatusCode)
}
