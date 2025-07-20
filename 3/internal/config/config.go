package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	defaultUrl  = "https://example.com"
	defaultJson = "json/data.json"
	defaultLogs = "logs/data.log"
	defaultRes = "result/result.txt"
)

type Config struct {
	URL  string
	Json string
	Logs string
	Res string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		return &Config{
			URL:  defaultUrl,
			Json: defaultJson,
			Logs: defaultLogs,
			Res: defaultRes,
		}
	}
	return &Config{
		URL:  os.Getenv("URL"),
		Json: os.Getenv("JSON"),
		Logs: os.Getenv("LOGS"),
		Res: os.Getenv("RES"),
	}

}
