package writer

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
)

func (f *fileWrite) WriteFile(name string, data []byte, perm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(name), 0750); err != nil {
		return err
	}
	return os.WriteFile(name, data, perm)
}
func (f *fileWrite) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(name))
}

func (f *fileWrite) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(filepath.Clean(name), flag, perm)
}

func (f *fileWrite) Write(file *os.File, b []byte) (n int, err error) {
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(chalk.Red("error closing file: %s"), err)
		}
	}()
	return file.Write(b)
}
