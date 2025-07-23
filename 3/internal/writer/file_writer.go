package writer

import (
	"fmt"
	"os"
)

type FileWriter struct {
	file *os.File
}

func NewFileWriter(filename string) (*FileWriter, error) {
	const op = "op.writer.NewFileWriter"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return &FileWriter{
		file: file,
	}, nil
}

func (w *FileWriter) WriteData(data string) error {
	const op = "op.writer.WriteData"
	if _, err := w.file.WriteString(data); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (w *FileWriter) Close() error {
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}
