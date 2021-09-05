package engine

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
)

const (
	FileModeOwnerRWX = 0644
)

// WriteFile thin wrapper to decouple dependency of write file
func (w *writer) WriteFile(name string, data []byte, perm fs.FileMode) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	return w.fs.WriteFile(name, data, perm)
}

func (w *writer) AppendFile(name string, data []byte) error {
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

func (w *writer) InjectIntoFile(name string, data []byte, inject inject) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	source, err := w.fs.ReadFile(name)
	if err != nil {
		log.Println("error reading file", err)
		return err
	}
	formatedOutput := injectIntoData(name, source, data, inject)
	if err := w.fs.WriteFile(name, []byte(formatedOutput), FileModeOwnerRWX); err != nil {
		log.Println("error appending data", err)
		return err
	}
	return nil
}

func injectIntoData(name string, source, data []byte, inject inject) []byte {
	var splitByMatcher []string

	switch inject.After {
	case true:
		splitByMatcher = strings.SplitAfter(string(source), inject.Matcher)
		if len(splitByMatcher) != 2 {
			log.Printf("injection token %s is not found in file %s", inject.Matcher, name)
			return source
		}
	default:
		idx := strings.LastIndex(string(source), inject.Matcher)
		if idx == -1 {
			log.Printf("injection token %s is not found in file %s", inject.Matcher, name)
			return source
		}
		splitByMatcher = []string{
			string(source[:(idx - 1)]),
			string(source[idx:]),
		}
	}

	formatedOutput := strings.Join([]string{
		splitByMatcher[0],
		string(data),
		splitByMatcher[1],
	}, "")
	return []byte(formatedOutput)
}

func isAfter(before string, after string) bool {
	return len(strings.TrimSpace(after)) > 0
}

func getMatcher(before, after string) string {
	if len(strings.TrimSpace(before)) > 0 {
		return before
	}
	return after
}

func setFileWriter(dryrun bool) fileReadWrite {
	if dryrun {
		return &fileLog{}
	}
	return &fileWrite{}
}

func (f *fileWrite) WriteFile(name string, data []byte, perm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(name), 0750); err != nil {
		return err
	}
	return os.WriteFile(name, data, perm)
}
func (f *fileWrite) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (f *fileWrite) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (f *fileWrite) Write(file *os.File, b []byte) (n int, err error) {
	defer file.Close()
	return file.Write(b)
}

func (f *fileLog) WriteFile(name string, data []byte, perm os.FileMode) error {
	log.Println(chalk.Green("logging to console:"), name, fmt.Sprintf("\n%s", string(data)))
	return nil
}
func (f *fileLog) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (f *fileLog) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (f *fileLog) Write(file *os.File, b []byte) (n int, err error) {
	log.Println(chalk.Green("logging to console:"), file.Name(), fmt.Sprintf("\n%s", string(b)))
	return 0, nil
}
