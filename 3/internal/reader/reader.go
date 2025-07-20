package reader

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonReader struct {
	filename string
}

func NewReader(fileName string) *JsonReader {
	return &JsonReader{
		filename: fileName,
	}
}

func (j *JsonReader) ReadFile() ([]int64, error) {
	const op = "op.reader.ReadFile"
	var mySlice []int64
	data, err := os.Open(j.filename)
	if err != nil {

		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer data.Close()

	if err = json.NewDecoder(data).Decode(&mySlice); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return mySlice, nil
}

// func (j *JsonReader) ReadFile() ([]int64, error) {
// 	const op = "op.reader.ReadFile"
// 	var mySlice []int64
// 	data, err := os.ReadFile(j.filename)
// 	if err != nil {

// 		return nil, fmt.Errorf("%s: %s", op, err)
// 	}
// 	if err = json.Unmarshal(data, &mySlice); err != nil {
// 		return nil, fmt.Errorf("%s: %s", op, err)
// 	}
// 	return mySlice, nil
// }
