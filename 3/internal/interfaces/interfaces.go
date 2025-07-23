package interfaces

import "net/http"

type DataReader interface {
	ReadData() ([]int64, error)
}
type DataWriter interface {
	WriteData(data string) error
	Close() error
}
type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Close() error
}
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}
type Calculator interface {
	Sum(numbers []int64) int64
}
