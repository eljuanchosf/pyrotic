package parser

import (
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/formats"
)

const (
	caseSnake  = "caseSnake"
	caseKebab  = "caseKebab"
	casePascal = "casePascal"
	caseLower  = "caseLower"
	caseTitle  = "caseTitle"
	caseCamel  = "caseCamel"
)

var (
	defaultFuncs = template.FuncMap{
		caseSnake:  formats.CaseSnake,
		caseKebab:  formats.CaseKebab,
		casePascal: formats.CasePascal,
		caseLower:  strings.ToLower,
		caseTitle:  strings.ToTitle,
		caseCamel:  formats.CaseCamel,
	}
)

func New(dirPath string, fileSuffix string) (TmplEngine, error) {
	tmp := template.New("root").Funcs(defaultFuncs)
	tmp, err := withTemplates(tmp, fileSuffix, dirPath)
	if err != nil {
		log.Println("error loading templates ", err)
		return TmplEngine{}, err
	}
	return TmplEngine{
		root:  tmp,
		funcs: defaultFuncs,
	}, nil
}

func (te *TmplEngine) Parse(data TemplateData) ([]TemplateData, error) {
	result := []TemplateData{}
	tmp := te.root.Templates()
	for _, t := range tmp {
		newData, err := parse(t.Root.String(), data, te.funcs)
		if err != nil {
			log.Println("error parsing template ", err)
			return result, err
		}

		formattedOut, err := format.Source(newData.Output)
		if err != nil {
			log.Println("error formatting ", err)
			return result, err
		}
		result = append(result, TemplateData{
			Name:   newData.Name,
			Append: newData.Append,
			Inject: newData.Inject,
			Before: newData.Before,
			After:  newData.After,
			To:     newData.To,
			Output: formattedOut,
			Meta:   newData.Meta,
		})
	}
	return result, nil
}

// withTemplates - load templates by file path
func withTemplates(root *template.Template, fileSuffix string, dirPath string) (*template.Template, error) {

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return root, err
	}
	var allFiles []string
	for _, file := range files {
		fileLocation := filepath.Join(dirPath, file.Name())
		if strings.HasSuffix(file.Name(), fileSuffix) {
			log.Println(chalk.Green("loading template: "), fileLocation)
			allFiles = append(allFiles, fileLocation)
		}
	}
	root, err = root.ParseFiles(allFiles...)
	if err != nil {
		return root, err
	}
	return root, nil
}
