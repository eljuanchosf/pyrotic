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
	tmp, err := withTemplates(fileSuffix, dirPath)
	if err != nil {
		log.Println("error loading templates ", err)
		return TmplEngine{}, err
	}
	return TmplEngine{
		templates: tmp,
		funcs:     defaultFuncs,
	}, nil
}

func (te *TmplEngine) Parse(data TemplateData) ([]TemplateData, error) {
	result := []TemplateData{}
	for _, t := range te.templates {
		newData, err := parse(t, data, te.funcs)
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
func withTemplates(fileSuffix string, dirPath string) ([]string, error) {
	var rootTemplates []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return rootTemplates, err
	}

	for _, file := range files {
		fileLocation := filepath.Join(dirPath, file.Name())
		if strings.HasSuffix(file.Name(), fileSuffix) {
			log.Println(chalk.Green("loading template: "), fileLocation)
			data, err := os.ReadFile(fileLocation)
			if err != nil {
				log.Printf("error reading file %s", fileLocation)
				return rootTemplates, err
			}
			rootTemplates = append(rootTemplates, string(data))
		}
	}
	return rootTemplates, nil
}
