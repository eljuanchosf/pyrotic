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
)

func Parse(raw string, data Data) []byte {
	meta, output := extractMeta(raw)
	hydratedData := hydrateData(meta, data)
	foo, _ := template.New("root").Parse(output)

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)
	_ = foo.Execute(wr, data)
	if err := wr.Flush(); err != nil {
		log.Println("error flushing writer ", err)
		return buf.Bytes(), err
	}
	return output
}

func hydrateData(meta []string, data Data) Data {
	result := Data{
		Name: data.Name,
	}
	tmp := map[string]interface{}{}
	for _, item := range meta {
		if strings.Contains(item, fieldTo) {
			list := strings.Split(strings.TrimSpace(item), ":")
			log.Println(list)
			if len(list) == 2 {
				result.To = list[1]
			}
			continue
		}
		list := strings.Split(strings.TrimSpace(item), ":")
		if len(list) == 2 {
			tmp[list[0]] = list[1]
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
