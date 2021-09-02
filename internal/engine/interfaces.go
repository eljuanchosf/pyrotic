package engine

import "os"

type fileWriter interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
}
