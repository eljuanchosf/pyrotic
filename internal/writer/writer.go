package writer

import (
	"io/fs"
	"log"
	"os"
	"sync"
)

const (
	FileModeOwnerRWX = 0644
)

func New(dryrun bool) Write {
	return Write{
		mx: sync.RWMutex{},
		fs: setFileWriter(dryrun),
	}
}

// WriteFile thin wrapper to decouple dependency of write file
func (w *Write) WriteFile(name string, data []byte, perm fs.FileMode) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	return w.fs.WriteFile(name, data, perm)
}

func (w *Write) AppendFile(name string, data []byte) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	file, err := w.fs.OpenFile(name, os.O_APPEND|os.O_WRONLY, FileModeOwnerRWX)
	if err != nil {
		log.Println("error opening file", err)
		return err
	}
	if _, err := w.fs.Write(file, data); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}

func (w *Write) InjectIntoFile(name string, data []byte, inject Inject) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	source, err := w.fs.ReadFile(name)
	if err != nil {
		log.Println("error reading file", err)
		return err
	}
	formatedOutput := mergeOutputs(name, source, data, inject)
	if err := w.fs.WriteFile(name, []byte(formatedOutput), FileModeOwnerRWX); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}

func setFileWriter(dryrun bool) fileReadWrite {
	if dryrun {
		return &fileLog{}
	}
	return &fileWrite{}
}
