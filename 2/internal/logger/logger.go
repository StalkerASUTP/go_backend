package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	file    *os.File
	*slog.Logger
}

func NewLogger(fileName string) (*Logger, error) {
	const op = "op.logger.NewLogger"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return &Logger{
		file: logFile,
		Logger: slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}, nil

}

func (l *Logger) CloseFile() {
	l.file.Close()
}
