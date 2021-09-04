package engine

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

const (
	FileModeOwnerRWX = 0644
)

// WriteFile thin wrapper to decouple dependency of write file
func (w *writer) WriteFile(name string, data []byte, perm fs.FileMode) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	return os.WriteFile(name, data, perm)
}

func (w *writer) AppendFile(name string, data []byte) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, FileModeOwnerRWX)
	if err != nil {
		log.Println("error opening file", err)
		return err
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}

func (w *writer) Inject(name string, data []byte, matcher string) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	file, err := os.ReadFile(name)
	if err != nil {
		log.Println("error reading file", err)
		return err
	}
	idx := strings.LastIndex(string(file), matcher)
	log.Println("file ", string(file[idx]))
	if err := os.WriteFile(name, data, FileModeOwnerRWX); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}
