package engine

import (
	"log"
	"sync"

	"github.com/code-gorilla-au/pyrotic/internal/parser"
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
		fwr:    setWriter(dryrun),
	}, nil
}

// Generate - generates code from
func (c *Core) Generate(data Data) error {

	parsedOutput, err := c.parser.Parse(parser.TemplateData{
		Name: data.Name,
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
			if err := c.fwr.InjectIntoFile(item.To, item.Output, inject{
				After:   isAfter(item.Before, item.After),
				Matcher: getMatcher(item.Before, item.After),
			}); err != nil {
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

func setWriter(dryrun bool) writer {
	return writer{
		mx: sync.RWMutex{},
		fs: setFileWriter(dryrun),
	}
}
