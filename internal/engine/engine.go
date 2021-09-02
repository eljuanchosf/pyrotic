package engine

import (
	"go/format"
	"log"
)

func New(dirPath string, fileSuffix string) (Core, error) {
	tmpl, err := newTmplEngine(dirPath, fileSuffix)
	if err != nil {
		return Core{}, err
	}
	return Core{
		parser: tmpl,
		fwr:    writer{},
	}, nil
}

// Generate - generates code from
func (c *Core) Generate(data Data) error {
	tmp := c.parser.root.Templates()
	for _, t := range tmp {
		newData, rawOutput, err := parse(t.Root.String(), data, c.parser.funcs)
		if err != nil {
			log.Println("error parsing template ", err)
			return err
		}

		formattedOut, err := format.Source(rawOutput)
		if err != nil {
			log.Println("error formatting ", err)
			return err
		}

		if err := c.fwr.WriteFile(newData.To, formattedOut, 0750); err != nil {
			log.Println("error writing to file ", err)
			return err
		}
	}
	return nil
}
