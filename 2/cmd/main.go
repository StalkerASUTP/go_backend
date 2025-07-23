package main

import (
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/factory"
	"go-backend/internal/service"
	"os"
)

func main() {

	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Создаем фабрику
	f := factory.NewFactory()

	// Создаем компоненты через фабрику
	logger, err := f.CreateLogger(cfg.Logs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	dataReader := f.CreateDataReader(cfg.Json)

	calculator := f.CreateCalculator()
	httpClient := f.CreateHTTPClient()

	// Создаем сервис
	processor := service.NewDataProcessor(dataReader, logger, calculator, httpClient)

	// Обрабатываем данные
	if err := processor.ProcessData(); err != nil {
		logger.Error("data processing failed", "error", err)
		os.Exit(1)
	}
	// Выполняем HTTP запрос
	if err := processor.ProcessHTTPRequest(cfg.URL); err != nil {
		logger.Error("HTTP request processing failed", "error", err)
		os.Exit(1)
	}
	logger.Info("application completed successfully")
}
