package main

import (
	"flag"
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/factory"
	"go-backend/internal/service"
	"os"
)

func main() {
	// Парсим флаги
	mode := flag.String("mode", "json", "mode of reading numbers (json or stdin)")
	flag.Parse()
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

	dataReader, err := f.CreateDataReader(*mode, cfg.Json)
	if err != nil {
		logger.Error("failed to create data reader", "mode", *mode, "error", err)
		os.Exit(1)
	}

	dataWriter, err := f.CreateDataWriter(cfg.Res)
	if err != nil {
		logger.Error("failed to create data writer", "error", err)
		os.Exit(1)
	}
	defer dataWriter.Close()

	calculator := f.CreateCalculator()
	httpClient := f.CreateHTTPClient()

	// Создаем сервис
	processor := service.NewDataProcessor(dataReader, dataWriter, logger, calculator, httpClient)

	logger.Info("application started", "mode", *mode)

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