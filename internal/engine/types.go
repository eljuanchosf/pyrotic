package engine

import "text/template"

type Core struct {
	root *template.Template
}

type Data struct {
	Name string
	To   string
	Meta map[string]interface{}
}
