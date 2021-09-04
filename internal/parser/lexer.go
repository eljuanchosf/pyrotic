package parser

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
	"text/template"
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
		tokens := strings.Split(strings.TrimSpace(item), tokenColon)
		if len(tokens) != 2 {
			log.Println("malformed template data for ", data.Name)
			break
		}
		switch strings.TrimSpace(tokens[0]) {
		case fieldTo:
			result.To = strings.TrimSpace(tokens[1])
		case fieldAfter:
			result.After = strings.TrimSpace(tokens[1])
		case fieldBefore:
			result.Before = strings.TrimSpace(tokens[1])
		case fieldAppend:
			stringAppend := strings.TrimSpace(tokens[1])
			append, err := strconv.ParseBool(stringAppend)
			if err != nil {
				log.Println("error parsing bool", err)
			}
			result.Append = append
		case fieldInject:
			stringAppend := strings.TrimSpace(tokens[1])
			inject, err := strconv.ParseBool(stringAppend)
			if err != nil {
				log.Println("error parsing bool", err)
			}
			result.Inject = inject
		default:
			key := strings.TrimSpace(tokens[0])
			tmp[key] = strings.TrimSpace(tokens[1])
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
