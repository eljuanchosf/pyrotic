package engine

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"text/template"
)

const (
	fieldTo = "to"
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
func parse(tmpl *template.Template, data Data) (Data, []byte, error) {
	raw := tmpl.Root.String()
	meta, stringOutput := extractMeta(raw)

	hydratedData, err := generateMetaData(meta, data)
	if err != nil {
		return hydratedData, nil, err
	}
	output, err := generateTemplate(string(stringOutput), hydratedData)
	if err != nil {
		return hydratedData, nil, err
	}
	return hydratedData, output, nil
}

func generateMetaData(meta []string, data Data) (Data, error) {
	parsedMeta := []string{}
	for _, item := range meta {
		t, err := template.New("meta").Parse(item)
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

func generateTemplate(output string, data Data) ([]byte, error) {
	tmpl, err := template.New("root").Parse(output)
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

func hydrateData(meta []string, data Data) Data {
	result := Data{
		Name: data.Name,
	}
	tmp := map[string]interface{}{}
	for _, item := range meta {
		if strings.Contains(item, fieldTo) {
			list := strings.Split(strings.TrimSpace(item), tokenColon)
			if len(list) == 2 {
				result.To = strings.TrimSpace(list[1])
			}
			continue
		}
		list := strings.Split(strings.TrimSpace(item), tokenColon)
		if len(list) == 2 {
			tmp[list[0]] = strings.TrimSpace(list[1])
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
