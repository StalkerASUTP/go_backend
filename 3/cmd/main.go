package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/logger"
	"go-backend/internal/reader"
	"go-backend/internal/request"
	"go-backend/internal/writer"
	"os"
	"strconv"
	"strings"
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
	mode := flag.String("mode", "json", "mode of reading numbers")
	flag.Parse()
	logger, err := logger.NewLogger(cfg.Logs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to inisialise logger: %s\n", err)
		return
	}

	defer logger.CloseFile()
	logger.Info("logger inisialised")
	logger.Info("mode of reading numbers", "mode", *mode)
	logger.Info("config loaded")
	writer, err := writer.NewWriter(cfg.Res)
	if err != nil {
		logger.Error("failed to inisialise writer", "error", err)
		return
	}
	defer writer.CloseFile()
	var slice []int64
	if *mode == "json" {
		slice, err = reader.NewReader(cfg.Json).ReadFile()
		if err != nil {
			logger.Error("failed to read json file", "error", err)
			return
		}
		logger.Info("Numbers have been read from JSON file")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter a number (or just Enter to finish): ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				break
			}
			number, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Please enter a valid number.")
				continue
			}
			slice = append(slice, number)
		}
	}

	logger.Info("Total sum of numbers is", "sum", sum(slice))
	resStr := fmt.Sprintf("Total sum of numbers is: %d\n", sum(slice))
	err = writer.NewRecord(resStr)
	if err != nil {
		logger.Error("failed to add new record", "result", cfg.Res)
	}
	fmt.Fprint(os.Stdout, resStr)
	res, err := request.NewRequest(cfg.URL).GetResponse()
	if err != nil {
		logger.Error("failed to get response", "url", cfg.URL)
		return
	}
	logger.Info("The response is getted", "host", res.Request.Host, "status code", res.StatusCode)
	resStr = fmt.Sprintf("The response is getted. Host: %s, status code: %d\n", res.Request.Host, res.StatusCode)
	err = writer.NewRecord(resStr)
	if err != nil {
		logger.Error("failed to add new record", "result", cfg.Res)
	}
	fmt.Fprint(os.Stdout, resStr)
}
