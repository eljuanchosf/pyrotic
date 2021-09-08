package writer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
)

func (f *fileLog) WriteFile(name string, data []byte, perm os.FileMode) error {
	log.Println(chalk.Green("logging to console:"), name, fmt.Sprintf("\n%s", string(data)))
	return nil
}
func (f *fileLog) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (f *fileLog) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(filepath.Clean(name), flag, perm)
}

func (f *fileLog) Write(file *os.File, b []byte) (n int, err error) {
	log.Println(chalk.Green("logging to console:"), file.Name(), fmt.Sprintf("\n%s", string(b)))
	return 0, nil
}
