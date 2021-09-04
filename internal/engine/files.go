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

func (w *writer) Inject(name string, data []byte, inject inject) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	source, err := os.ReadFile(name)
	if err != nil {
		log.Println("error reading file", err)
		return err
	}
	splitByMatcher := strings.SplitAfter(string(source), inject.Matcher)
	if len(splitByMatcher) != 2 {
		log.Printf("injection token %s is not found in file %s", inject.Matcher, name)
		return nil
	}
	formatedOutput := injectIntoData(name, source, data, inject)
	if err := os.WriteFile(name, []byte(formatedOutput), FileModeOwnerRWX); err != nil {
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
