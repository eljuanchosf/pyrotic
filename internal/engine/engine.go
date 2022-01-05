package engine

import (
	"log"
	"strings"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/parser"
	"github.com/code-gorilla-au/pyrotic/internal/writer"
)

const (
	metaDelimiter         = ","
	metaKeyValueDelimiter = "="
)

func New(dryrun bool, dirPath string, sharedPath string, fileSuffix string) (Core, error) {
	if dryrun {
		log.Println(chalk.Cyan("DRYRUN MODE"))
	}
	tmpl, err := parser.New(dirPath, sharedPath, fileSuffix)
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
		ParseData: parser.ParseData{
			Meta: generateMeta(data.MetaArgs),
		},
	})
	if err != nil {
		return err
	}

	for _, item := range parsedOutput {
		switch item.ParseData.Action {
		case parser.ActionAppend:
			if err := c.fwr.AppendFile(item.To, item.Output); err != nil {
				log.Println("error appending file ", err)
				return err
			}
		case parser.ActionInject:
			if err := c.fwr.InjectIntoFile(item.To, item.Output, writer.Inject{
				Matcher: item.InjectMatcher,
				Clause:  writer.InjectClause(item.InjectClause),
			}); err != nil {
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

func generateMeta(meta string) map[string]string {
	result := map[string]string{}

	list := strings.Split(meta, metaDelimiter)
	for _, keyVal := range list {
		rawMeta := strings.Split(keyVal, strings.TrimSpace(metaKeyValueDelimiter))

		if len(rawMeta) == 0 {
			continue
		}
		key := strings.TrimSpace(rawMeta[0])
		value := strings.TrimSpace(rawMeta[1])
		result[key] = value
	}

	return result
}
