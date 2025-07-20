package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	filename string
}

func NewLogger(fileName string) *Logger {
	return &Logger{
		filename: fileName,
	}
}

func (l *Logger) InisialiseLog() (*slog.Logger, *os.File, error) {
	const op = "op.logger.InisialiseLog"
	logFile, err := os.OpenFile(l.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil,nil, fmt.Errorf("%s: %s", op, err)
	}
	return slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})), logFile, nil

}
