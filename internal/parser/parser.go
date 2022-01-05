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
	"github.com/gobuffalo/flect"
)

const (
	caseSnake   = "caseSnake"
	caseKebab   = "caseKebab"
	casePascal  = "casePascal"
	caseLower   = "caseLower"
	caseTitle   = "caseTitle"
	caseCamel   = "caseCamel"
	pluralise   = "pluralise"
	singularise = "singularise"
	ordinalize  = "ordinalize"
	titleize    = "titleize"
	humanize    = "humanize"
)

var (
	defaultFuncs = template.FuncMap{
		caseSnake:  formats.CaseSnake,
		caseKebab:  formats.CaseKebab,
		casePascal: formats.CasePascal,
		caseLower:  strings.ToLower,
		caseTitle:  strings.ToTitle,
		caseCamel:  formats.CaseCamel,
		// Inflections
		pluralise:   flect.Pluralize,
		singularise: flect.Singularize,
		ordinalize:  flect.Ordinalize,
		titleize:    flect.Titleize,
		humanize:    flect.Humanize,
	}
)

func New(dirPath string, fileSuffix string) (TemplateEngine, error) {
	tmp, err := withTemplates(fileSuffix, dirPath)
	if err != nil {
		log.Printf(chalk.Red("error loading templates: %s"), err)
		return TemplateEngine{}, err
	}
	return TemplateEngine{
		templates: tmp,
		funcs:     defaultFuncs,
	}, nil
}

func (te *TemplateEngine) Parse(data TemplateData) ([]TemplateData, error) {
	result := []TemplateData{}
	for _, t := range te.templates {
		newData, err := parse(t, data, te.funcs)
		if err != nil {
			log.Printf(chalk.Red("error parsing template: %s"), err)
			return result, err
		}

		formattedOut, err := format.Source(newData.Output)
		if err != nil {
			log.Printf(chalk.Red("error formatting: %s"), err)
			return result, err
		}
		newData.Output = formattedOut
		result = append(result, newData)
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
				log.Printf(chalk.Red("error reading file %s"), fileLocation)
				return rootTemplates, err
			}
			rootTemplates = append(rootTemplates, string(data))
		}
	}
	return rootTemplates, nil
}
