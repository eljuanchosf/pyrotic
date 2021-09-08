package writer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
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

func mergeOutputs(name string, source, output []byte, inject Inject) []byte {
	var splitByMatcher []string
	if !inject.Validate() {
		log.Printf("at least 1 injection clause must not be empty, before: [%s], after: [%s]", inject.Before, inject.After)
		return source
	}
	switch isAfter(inject.Before, inject.After) {
	case true:
		splitByMatcher = strings.SplitAfter(string(source), inject.After)
		if len(splitByMatcher) != 2 {
			log.Printf("injection token %s is not found in file %s", inject.After, name)
			return source
		}
	default:
		idx := strings.Index(string(source), inject.Before)
		if idx == -1 {
			log.Printf("injection token %s is not found in file %s", inject.Before, name)
			return source
		}
		splitByMatcher = []string{
			string(source[:(idx - 1)]),
			fmt.Sprintf("\n%s", string(source[idx:])),
		}
	}

	formatedOutput := strings.Join([]string{
		splitByMatcher[0],
		string(output),
		splitByMatcher[1],
	}, "")
	return []byte(formatedOutput)
}

func isAfter(before string, after string) bool {
	return len(strings.TrimSpace(after)) > 0
}

func setFileWriter(dryrun bool) fileReadWrite {
	if dryrun {
		return &fileLog{}
	}
	return &fileWrite{}
}

// Validate - one clause must be met
func (i *Inject) Validate() bool {
	return i.After != "" || i.Before != ""
}
