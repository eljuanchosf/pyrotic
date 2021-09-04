package parser

import (
	"bufio"
	"bytes"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

const (
	fieldTo     = "to"
	fieldAppend = "append"
)

const (
	tokenNewLine = "\n"
	tokenDash    = "---"
	tokenColon   = ":"
)

// parse - 2 stage parse for a template.
//
// stage 1: hydrate the data from the metadata within the "---" block of the template
//
// stage 2: parse and execute the template with the hydrated metadata
func parse(raw string, data TemplateData, funcs template.FuncMap) (TemplateData, error) {
	meta, stringOutput := extractMeta(raw)

	hydratedData, err := generateMetaData(meta, data, funcs)
	if err != nil {
		return hydratedData, err
	}
	output, err := generateTemplate(string(stringOutput), hydratedData, funcs)
	if err != nil {
		return hydratedData, err
	}
	hydratedData.Output = output
	return hydratedData, nil
}

func generateMetaData(meta []string, data TemplateData, funcs template.FuncMap) (TemplateData, error) {
	parsedMeta := []string{}
	for _, item := range meta {
		t, err := template.New("meta").Funcs(funcs).Parse(item)
		if err != nil {
			log.Println("error generating metadata ", err)
			return data, err
		}
		var buf bytes.Buffer
		wr := bufio.NewWriter(&buf)
		if t.Execute(wr, data); err != nil {
			log.Println("error executing template ", err)
			return data, err
		}
		if err := wr.Flush(); err != nil {
			log.Println("error flushing writer ", err)
			return data, err
		}
		parsedMeta = append(parsedMeta, buf.String())
	}
	return hydrateData(parsedMeta, data), nil
}

func generateTemplate(output string, data TemplateData, funcs template.FuncMap) ([]byte, error) {
	tmpl, err := template.New("root").Funcs(funcs).Parse(output)
	if err != nil {
		log.Println("error parsing output ", err)
		return nil, err
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)
	if err := tmpl.Execute(wr, data); err != nil {
		log.Println("error executing template ", err)
		return nil, err
	}
	if err := wr.Flush(); err != nil {
		log.Println("error flushing writer ", err)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

func hydrateData(meta []string, data TemplateData) TemplateData {
	result := TemplateData{
		Name: data.Name,
	}
	tmp := map[string]interface{}{}
	for _, item := range meta {
		switch {
		case strings.Contains(item, fieldTo):
			list := strings.Split(strings.TrimSpace(item), tokenColon)
			if len(list) == 2 {
				result.To = strings.TrimSpace(list[1])
			}
		case strings.Contains(item, fieldAppend):
			list := strings.Split(strings.TrimSpace(item), tokenColon)

			if len(list) == 2 {
				stringAppend := strings.TrimSpace(list[1])
				append, err := strconv.ParseBool(stringAppend)
				if err != nil {
					log.Println("error parsing bool", err)
				}
				result.Append = append
			}
		default:
			list := strings.Split(strings.TrimSpace(item), tokenColon)
			if len(list) == 2 {
				key := strings.TrimSpace(list[0])
				tmp[key] = strings.TrimSpace(list[1])
			}
		}
	}
	result.Meta = tmp
	return result
}

func extractMeta(output string) ([]string, string) {
	list := strings.Split(string(output), tokenNewLine)
	meta := []string{}
	count := 0
	for index, s := range list {
		if count == 2 {
			list = list[index:]
			break
		}

		if s == tokenDash {
			count++
			continue
		}
		meta = append(meta, s)
	}
	formattedOutput := strings.Join(list, tokenNewLine)
	return meta, formattedOutput
}
