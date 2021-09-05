package engine

import (
	"sync"

	"github.com/code-gorilla-au/pyrotic/internal/parser"
)

type Core struct {
	parser parser.TmplEngine
	fwr    writer
}

type Data struct {
	Name string
}

type writer struct {
	mx sync.RWMutex
	fs fileReadWrite
}

type fileWrite struct {
	DryRun bool
}

var _ fileReadWrite = (*fileWrite)(nil)

type fileLog struct{}

var _ fileReadWrite = (*fileLog)(nil)

type inject struct {
	After   bool
	Matcher string
}
