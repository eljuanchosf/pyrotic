// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package writer

import (
	"io/fs"
	"os"
	"sync"
)

// Ensure, that fileReadWriteMock does implement fileReadWrite.
// If this is not the case, regenerate this file with moq.
var _ fileReadWrite = &fileReadWriteMock{}

// fileReadWriteMock is a mock implementation of fileReadWrite.
//
// 	func TestSomethingThatUsesfileReadWrite(t *testing.T) {
//
// 		// make and configure a mocked fileReadWrite
// 		mockedfileReadWrite := &fileReadWriteMock{
// 			OpenFileFunc: func(name string, flag int, perm fs.FileMode) (*os.File, error) {
// 				panic("mock out the OpenFile method")
// 			},
// 			ReadFileFunc: func(name string) ([]byte, error) {
// 				panic("mock out the ReadFile method")
// 			},
// 			WriteFunc: func(file *os.File, b []byte) (int, error) {
// 				panic("mock out the Write method")
// 			},
// 			WriteFileFunc: func(name string, data []byte, perm fs.FileMode) error {
// 				panic("mock out the WriteFile method")
// 			},
// 		}
//
// 		// use mockedfileReadWrite in code that requires fileReadWrite
// 		// and then make assertions.
//
// 	}
type fileReadWriteMock struct {
	// OpenFileFunc mocks the OpenFile method.
	OpenFileFunc func(name string, flag int, perm fs.FileMode) (*os.File, error)

	// ReadFileFunc mocks the ReadFile method.
	ReadFileFunc func(name string) ([]byte, error)

	// WriteFunc mocks the Write method.
	WriteFunc func(file *os.File, b []byte) (int, error)

	// WriteFileFunc mocks the WriteFile method.
	WriteFileFunc func(name string, data []byte, perm fs.FileMode) error

	// calls tracks calls to the methods.
	calls struct {
		// OpenFile holds details about calls to the OpenFile method.
		OpenFile []struct {
			// Name is the name argument value.
			Name string
			// Flag is the flag argument value.
			Flag int
			// Perm is the perm argument value.
			Perm fs.FileMode
		}
		// ReadFile holds details about calls to the ReadFile method.
		ReadFile []struct {
			// Name is the name argument value.
			Name string
		}
		// Write holds details about calls to the Write method.
		Write []struct {
			// File is the file argument value.
			File *os.File
			// B is the b argument value.
			B []byte
		}
		// WriteFile holds details about calls to the WriteFile method.
		WriteFile []struct {
			// Name is the name argument value.
			Name string
			// Data is the data argument value.
			Data []byte
			// Perm is the perm argument value.
			Perm fs.FileMode
		}
	}
	lockOpenFile  sync.RWMutex
	lockReadFile  sync.RWMutex
	lockWrite     sync.RWMutex
	lockWriteFile sync.RWMutex
}

// OpenFile calls OpenFileFunc.
func (mock *fileReadWriteMock) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	callInfo := struct {
		Name string
		Flag int
		Perm fs.FileMode
	}{
		Name: name,
		Flag: flag,
		Perm: perm,
	}
	mock.lockOpenFile.Lock()
	mock.calls.OpenFile = append(mock.calls.OpenFile, callInfo)
	mock.lockOpenFile.Unlock()
	if mock.OpenFileFunc == nil {
		var (
			fileOut *os.File
			errOut  error
		)
		return fileOut, errOut
	}
	return mock.OpenFileFunc(name, flag, perm)
}

// OpenFileCalls gets all the calls that were made to OpenFile.
// Check the length with:
//     len(mockedfileReadWrite.OpenFileCalls())
func (mock *fileReadWriteMock) OpenFileCalls() []struct {
	Name string
	Flag int
	Perm fs.FileMode
} {
	var calls []struct {
		Name string
		Flag int
		Perm fs.FileMode
	}
	mock.lockOpenFile.RLock()
	calls = mock.calls.OpenFile
	mock.lockOpenFile.RUnlock()
	return calls
}

// ReadFile calls ReadFileFunc.
func (mock *fileReadWriteMock) ReadFile(name string) ([]byte, error) {
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockReadFile.Lock()
	mock.calls.ReadFile = append(mock.calls.ReadFile, callInfo)
	mock.lockReadFile.Unlock()
	if mock.ReadFileFunc == nil {
		var (
			bytesOut []byte
			errOut   error
		)
		return bytesOut, errOut
	}
	return mock.ReadFileFunc(name)
}

// ReadFileCalls gets all the calls that were made to ReadFile.
// Check the length with:
//     len(mockedfileReadWrite.ReadFileCalls())
func (mock *fileReadWriteMock) ReadFileCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockReadFile.RLock()
	calls = mock.calls.ReadFile
	mock.lockReadFile.RUnlock()
	return calls
}

// Write calls WriteFunc.
func (mock *fileReadWriteMock) Write(file *os.File, b []byte) (int, error) {
	callInfo := struct {
		File *os.File
		B    []byte
	}{
		File: file,
		B:    b,
	}
	mock.lockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	mock.lockWrite.Unlock()
	if mock.WriteFunc == nil {
		var (
			nOut   int
			errOut error
		)
		return nOut, errOut
	}
	return mock.WriteFunc(file, b)
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//     len(mockedfileReadWrite.WriteCalls())
func (mock *fileReadWriteMock) WriteCalls() []struct {
	File *os.File
	B    []byte
} {
	var calls []struct {
		File *os.File
		B    []byte
	}
	mock.lockWrite.RLock()
	calls = mock.calls.Write
	mock.lockWrite.RUnlock()
	return calls
}

// WriteFile calls WriteFileFunc.
func (mock *fileReadWriteMock) WriteFile(name string, data []byte, perm fs.FileMode) error {
	callInfo := struct {
		Name string
		Data []byte
		Perm fs.FileMode
	}{
		Name: name,
		Data: data,
		Perm: perm,
	}
	mock.lockWriteFile.Lock()
	mock.calls.WriteFile = append(mock.calls.WriteFile, callInfo)
	mock.lockWriteFile.Unlock()
	if mock.WriteFileFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.WriteFileFunc(name, data, perm)
}

// WriteFileCalls gets all the calls that were made to WriteFile.
// Check the length with:
//     len(mockedfileReadWrite.WriteFileCalls())
func (mock *fileReadWriteMock) WriteFileCalls() []struct {
	Name string
	Data []byte
	Perm fs.FileMode
} {
	var calls []struct {
		Name string
		Data []byte
		Perm fs.FileMode
	}
	mock.lockWriteFile.RLock()
	calls = mock.calls.WriteFile
	mock.lockWriteFile.RUnlock()
	return calls
}
