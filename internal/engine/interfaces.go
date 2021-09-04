package engine

import "os"

type fileWriter interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
	AppendFile(name string, data []byte) error
	Inject(name string, data []byte, inject inject) error
}
