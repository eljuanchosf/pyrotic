package engine

import (
	"io/fs"
	"log"
	"os"
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
	fileData, err := os.ReadFile(name)
	if err != nil {
		log.Println("error reading file", err)
		return err
	}
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening file", err)
		return err
	}
	defer file.Close()
	sourceData := append(fileData, data...)
	if _, err := file.Write(sourceData); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}
