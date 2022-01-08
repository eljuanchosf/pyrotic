package parser

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
	"text/template"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
)

const (
	fieldTo     = "to"
	fieldAppend = "append"
	fieldInject = "inject"
	fieldAfter  = "after"
	fieldBefore = "before"
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
func parse(tmplName, tmpl string, data TemplateData, funcs template.FuncMap, sharedTmpl map[string]string) (TemplateData, error) {
	meta, stringOutput := extractMetaDataFromTemplate(tmpl)

	hydratedData, err := generateParseData(meta, data, funcs)
	if err != nil {
		return hydratedData, err
	}
	output, err := generateTemplate(tmplName, string(stringOutput), hydratedData, funcs, sharedTmpl)
	if err != nil {
		return hydratedData, err
	}
	hydratedData.Output = output
	return hydratedData, nil
}

func generateParseData(meta []string, data TemplateData, funcs template.FuncMap) (TemplateData, error) {
	parsedMeta := []string{}

	for _, item := range meta {
		var buf bytes.Buffer
		wr := bufio.NewWriter(&buf)
		t, err := template.New("meta").Funcs(funcs).Parse(item)
		if err != nil {
			log.Printf(chalk.Red("error generating metadata %s"), err)
			return data, err
		}

		if err := t.Execute(wr, data); err != nil {
			log.Printf(chalk.Red(err.Error()), "\n", item)
			return data, err
		}

		if err := wr.Flush(); err != nil {
			log.Printf(chalk.Red("error flushing writer %s"), err)
			return data, err
		}

		parsedMeta = append(parsedMeta, buf.String())
	}

	return hydrateData(parsedMeta, data)

}

func generateTemplate(tmplName, tmplOutput string, data TemplateData, funcs template.FuncMap, sharedTmpl map[string]string) ([]byte, error) {
	tmpl, err := template.New(tmplName).Funcs(funcs).Parse(tmplOutput)
	if err != nil {
		log.Printf(chalk.Red("error parsing output: %s"), err)
		return nil, err
	}

	for sharedTmplName, tmplOutput := range sharedTmpl {
		// we don't mind if this fails
		tmpl.New(sharedTmplName).Funcs(funcs).Parse(tmplOutput)
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)
	if err := tmpl.Execute(wr, data); err != nil {
		log.Println(chalk.Red(err.Error()), "\n", tmplOutput)
		return nil, err
	}
	if err := wr.Flush(); err != nil {
		log.Printf(chalk.Red("error flushing writer: %s"), err)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

func hydrateData(meta []string, data TemplateData) (TemplateData, error) {
	result := TemplateData{
		Name:      data.Name,
		ParseData: data.ParseData,
	}
	result.ParseData.Action = ActionCreate

	tmp := map[string]string{}
	for _, item := range meta {
		tokens := strings.Split(strings.TrimSpace(item), tokenColon)
		if len(tokens) != 2 {
			return result, ErrMalformedTemplate
		}

		switch strings.TrimSpace(tokens[0]) {
		case fieldTo:
			result.To = strings.TrimSpace(tokens[1])
		case fieldAfter:
			result.ParseData.InjectClause = InjectAfter
			result.ParseData.InjectMatcher = strings.TrimSpace(tokens[1])
		case fieldBefore:
			result.ParseData.InjectClause = InjectBefore
			result.ParseData.InjectMatcher = strings.TrimSpace(tokens[1])
		case fieldAppend:
			result.ParseData.Action = ActionAppend
			stringAppend := strings.TrimSpace(tokens[1])
			if _, err := strconv.ParseBool(stringAppend); err != nil {
				return result, ErrParsingBool
			}
		case fieldInject:
			result.ParseData.Action = ActionInject
			stringAppend := strings.TrimSpace(tokens[1])
			if _, err := strconv.ParseBool(stringAppend); err != nil {
				return result, ErrParsingBool
			}
		default:
			key := strings.TrimSpace(tokens[0])
			tmp[key] = strings.TrimSpace(tokens[1])
		}
	}

	// this will override any values pre-defined in the template,
	// this is indtended so you are able to have "sane defaults" as well as override via cmd
	for key, value := range data.Meta {
		tmp[key] = value
	}

	result.Meta = tmp
	return result, nil
}

func extractMetaDataFromTemplate(template string) ([]string, string) {
	rawOut := strings.Split(template, tokenNewLine)
	meta := []string{}
	output := []string{}
	count := 0
	for index, s := range rawOut {
		trimmed := strings.TrimSpace(s)
		if count == 2 {
			output = rawOut[index:]
			break
		}

		if trimmed == tokenDash {
			count++
			continue
		}
		if count >= 1 {
			meta = append(meta, trimmed)
		}
	}
	return meta, strings.Join(output, tokenNewLine)
}
