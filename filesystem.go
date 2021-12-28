package mockingdemo

import "fmt"

type FileSystem struct {
	conn FileConnector
}

func NewFS() FileSystem {
	fs := FileSystem{conn: NewFileConnector()}
	return fs
}

func (fs *FileSystem) PrintFileContent(fileName string) error {
	fileContent, err := fs.conn.ReadFromFile(fileName)

	if err != nil {
		return err
	}

	fmt.Println(string(fileContent))
	return nil
}
