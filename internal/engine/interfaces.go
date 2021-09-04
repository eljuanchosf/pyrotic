package engine

import (
	"io/fs"
	"os"
)

type fileWriter interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
	AppendFile(name string, data []byte) error
	Inject(name string, data []byte, inject inject) error
}

type files interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
	OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error)
	Write(b []byte) (n int, err error)
}
