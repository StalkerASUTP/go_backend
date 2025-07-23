package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type StructuredLogger struct {
	file   *os.File
	logger *slog.Logger
}

func NewStructuredLogger(filename string) (*StructuredLogger, error) {
	const op = "op.logger.NewStructuredLogger"
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	logger := slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return &StructuredLogger{
		file:   logFile,
		logger: logger,
	}, nil
}

func (l *StructuredLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}
func (l *StructuredLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
func (l *StructuredLogger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
