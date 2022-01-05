package parser

import "text/template"

type TemplateEngine struct {
	templates       []string
	sharedTemplates map[string]string
	funcs           template.FuncMap
}

type TemplateData struct {
	Name   string
	To     string
	Output []byte
	ParseData
}

type ParseActions string

const (
	ActionCreate ParseActions = "Create"
	ActionAppend ParseActions = "Append"
	ActionInject ParseActions = "Inject"
)

type InjectClause string

const (
	InjectBefore InjectClause = "Before"
	InjectAfter  InjectClause = "After"
)

type ParseData struct {
	Action         ParseActions
	InjectClause   InjectClause
	InjectMatcher  string
	SharedTemplate string
	Meta           map[string]string
}
