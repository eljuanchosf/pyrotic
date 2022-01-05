package writer

import (
	"io/fs"
	"log"
	"os"
	"sync"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
)

const (
	FileModeOwnerRWX = 0644
)

// New - create a new writer
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

// AppendFile - append data to the end of file
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

// InjectIntoFile - inject before, or after a matcher for a source file.
// If the matcher can't be found, don't do anything to the file
func (w *Write) InjectIntoFile(name string, data []byte, inject Inject) error {
	w.mx.Lock()
	defer w.mx.Unlock()
	source, err := w.fs.ReadFile(name)
	if err != nil {
		log.Printf(chalk.Red("error injecting data: %s"), err)
		return err
	}
	formatedOutput, err := mergeInjection(source, data, inject)
	if err != nil {
		log.Printf(chalk.Red("%s: file [%s], matcher: [%s], clause [%s]"), err, name, inject.Matcher, inject.Clause)
		return err
	}
	if err := w.fs.WriteFile(name, []byte(formatedOutput), FileModeOwnerRWX); err != nil {
		log.Println(chalk.Red("error writing to file"), err)
		return err
	}
	return nil
}

// setFileWriter - return a writer based on the dry run flag.
// If the dry fun flag is true, return a writer that logs to stdout,
// otherwise return a file writer.
func setFileWriter(dryrun bool) fileReadWrite {
	if dryrun {
		return &fileLog{}
	}
	return &fileWrite{}
}
