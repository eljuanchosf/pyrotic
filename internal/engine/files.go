package engine

import (
	"io/fs"
	"os"
)

// WriteFile thin wrapper to decouple dependency of write file
func (w writer) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}
