package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	defaultUrl  = "https://example.com"
	defaultJson = "json/data.json"
	defaultLogs = "logs/data.log"
)

type Config struct {
	URL  string
	Json string
	Logs string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	url := os.Getenv("URL")
	if url == "" {
		url = defaultUrl
	}
	json := os.Getenv("JSON")
	if json == "" {
		json = defaultJson
	}
	logs := os.Getenv("LOGS")
	if logs == "" {
		logs = defaultLogs
	}
	return &Config{
		URL:  url,
		Json: json,
		Logs: logs,
	}

}
