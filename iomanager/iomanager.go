package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data map[string]float64, key string) error
}
