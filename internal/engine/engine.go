package engine

import (
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/formats"
)

func New(dirPath string, fileSuffix string) (Core, error) {
	tmp, err := withTemplates(template.New("root"), fileSuffix, dirPath)
	if err != nil {
		return Core{}, err
	}
	tmp = withFuncs(tmp)
	return Core{
		root: tmp,
		fwr:  writer{},
	}, nil
}

func (c *Core) Generate(data Data) error {

	tmp := c.root.Templates()
	for _, t := range tmp {
		rawString := t.Root.String()
		rawOutput, err := parse(rawString, data)
		if err != nil {
			log.Println("error parsing template ", err)
			return err
		}
		formattedOut, err := format.Source(rawOutput)
		if err != nil {
			log.Println("error formatting ", err)
			return err
		}
		log.Println("to ", data.To)
		if err := c.fwr.WriteFile(data.To, formattedOut, 0600); err != nil {
			log.Println("error writing to file ", err)
			return err
		}
	}
	return nil
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
	caseSnake  = "caseSnake"
	caseKebab  = "caseKebab"
	casePascal = "casePascal"
)

func withFuncs(root *template.Template) *template.Template {
	return root.Funcs(template.FuncMap{
		caseSnake:  formats.CaseSnake,
		caseKebab:  formats.CaseKebab,
		casePascal: formats.CasePascal,
	})
}
