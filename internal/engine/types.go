package engine

import "github.com/code-gorilla-au/pyrotic/internal/parser"

type Core struct {
	parser parser.TmplEngine
	fwr    fileWriter
}

type Data struct {
	Name string
	To   string
	Meta map[string]interface{}
}

type writer struct {
}
