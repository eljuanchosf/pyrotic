package engine

type Core struct {
	parser TmplEngine
	fwr    fileWriter
}

type Data struct {
	Name string
	To   string
	Meta map[string]interface{}
}

type writer struct {
}
