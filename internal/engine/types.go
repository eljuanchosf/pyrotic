package engine

import (
	"github.com/code-gorilla-au/pyrotic/internal/parser"
	"github.com/code-gorilla-au/pyrotic/internal/writer"
)

type Core struct {
	parser parser.TmplEngine
	fwr    writer.Write
}

type Data struct {
	Name string
}
