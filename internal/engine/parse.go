package engine

import (
	"log"
	"strings"
)

func parse(output []byte) string {
	list := strings.Split(string(output), "\n")
	meta := []string{}
	count := 0
	for index, s := range list {
		if count == 2 {
			list = list[index:]
			break
		}

		if s == "---" {
			count++
			continue
		}
		meta = append(meta, s)
	}
	log.Println("meta", hydrateData(meta, Data{}))
	return strings.Join(list, "\n")
}

func hydrateData(meta []string, data Data) Data {
	var to string
	tmp := map[string]interface{}{}
	for _, item := range meta {
		if strings.Contains(item, "to") {
			list := strings.Split(strings.TrimSpace(item), ":")
			if len(list) == 2 {
				to = list[1]
			}
			continue
		}
		list := strings.Split(strings.TrimSpace(item), ":")
		if len(list) == 2 {
			tmp[list[0]] = list[1]
		}
	}
	return Data{
		Name: data.Name,
		To:   to,
		Meta: tmp,
	}
}
