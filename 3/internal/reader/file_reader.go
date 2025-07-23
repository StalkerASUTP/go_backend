package reader

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileReader struct {
	filename string
}

func NewFileReader(fileName string) *FileReader {
	return &FileReader{
		filename: fileName,
	}
}

func (j *FileReader) ReadData() ([]int64, error) {
	const op = "op.FileReader.ReadData"
	file, err := os.Open(j.filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer file.Close()
	var numbers []int64
	if err = json.NewDecoder(file).Decode(&numbers); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return numbers, nil
}

// более простой варинт для работы с небольшими json файлами

// func (j *FileReader) ReadData() ([]int64, error) {
// 	const op = "op.reader.ReadData"
// 	data, err := os.ReadFile(j.filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %s", op, err)
// 	}
// 	var numbers []int64
// 	if err = json.Unmarshal(data, &numbers); err != nil {
// 		return nil, fmt.Errorf("%s: %s", op, err)
// 	}
// 	return numbers, nil
// }
