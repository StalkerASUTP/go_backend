package writer

import (
	"fmt"
	"os"
)

type Writer struct {
	file *os.File
}

func NewWriter(fileName string) (*Writer, error) {
	const op = "op.writer.NewWriter"
	resFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return &Writer{
		file: resFile,
	}, nil
}

func (r *Writer) NewRecord(record string) error {
	const op = "op.writer.NewRecord"
	_, err := r.file.WriteString(record)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Writer) CloseFile() {
	r.file.Close()
}
