package engine

import "text/template"

type Core struct {
	root      *template.Template
	meta      *template.Template
	fwr       fileWriter
	tmplFuncs template.FuncMap
}

type Data struct {
	Name string
	To   string
	Meta map[string]interface{}
}

type writer struct {
}
