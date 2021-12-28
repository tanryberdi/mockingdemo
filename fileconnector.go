package mockingdemo

import "os"

type FileConnector interface {
	ReadFromFile(name string) ([]byte, error)
}

type fileConnector struct{}

func NewFileConnector() FileConnector {
	return &fileConnector{}
}

func (f *fileConnector) ReadFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, err
}
