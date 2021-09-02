package engine

import "text/template"

type Core struct {
	root *template.Template
	fwr  fileWriter
}

type Data struct {
	Name string
	To   string
	Meta map[string]interface{}
}

type writer struct {
}
