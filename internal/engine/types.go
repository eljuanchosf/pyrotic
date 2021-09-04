package engine

import (
	"sync"

	"github.com/code-gorilla-au/pyrotic/internal/parser"
)

type Core struct {
	parser parser.TmplEngine
	fwr    fileWriter
}

type Data struct {
	Name string
}

type writer struct {
	mx sync.RWMutex
}
