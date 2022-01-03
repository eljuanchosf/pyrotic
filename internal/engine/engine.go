package engine

import (
	"log"
	"strings"

	"github.com/code-gorilla-au/pyrotic/internal/parser"
	"github.com/code-gorilla-au/pyrotic/internal/writer"
)

const (
	metaDelimiter         = ","
	metaKeyValueDelimiter = "="
)

func New(dryrun bool, dirPath string, fileSuffix string) (Core, error) {
	if dryrun {
		log.Println("DRYRUN MODE")
	}
	tmpl, err := parser.New(dirPath, fileSuffix)
	if err != nil {
		return Core{}, err
	}
	return Core{
		parser: tmpl,
		fwr:    writer.New(dryrun),
	}, nil
}

// Generate - generates code from
func (c *Core) Generate(data Data) error {

	parsedOutput, err := c.parser.Parse(parser.TemplateData{
		Name: data.Name,
		Meta: generateMeta(data.Meta),
	})
	if err != nil {
		return err
	}

	for _, item := range parsedOutput {
		switch {
		case item.Append:
			if err := c.fwr.AppendFile(item.To, item.Output); err != nil {
				log.Println("error appending file ", err)
				return err
			}
		case item.Inject:
			if err := c.fwr.InjectIntoFile(item.To, item.Output, generateInject(item.Before, item.After)); err != nil {
				log.Println("error appending file ", err)
				return err
			}
		default:
			if err := c.fwr.WriteFile(item.To, item.Output, 0750); err != nil {
				log.Println("error writing to file ", err)
				return err
			}
		}

	}

	return nil
}

func generateInject(before, after string) writer.Inject {
	if len(before) > 0 {
		return writer.Inject{
			Matcher: before,
			Clause:  writer.InjectBefore,
		}
	}
	return writer.Inject{
		Matcher: after,
		Clause:  writer.InjectAfter,
	}
}

func generateMeta(meta string) map[string]string {
	result := map[string]string{}

	list := strings.Split(meta, metaDelimiter)
	for _, keyVal := range list {
		rawMeta := strings.Split(keyVal, strings.TrimSpace(metaKeyValueDelimiter))
		if len(rawMeta) != 0 {
			continue
		}
		result[rawMeta[0]] = rawMeta[1]
	}
	return result
}
