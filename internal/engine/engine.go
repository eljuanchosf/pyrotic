package engine

import (
	"bufio"
	"bytes"
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
	}, nil
}

func (c *Core) Generate(data Data) error {

	tmp := c.root.Templates()
	for _, t := range tmp {
		raw := t.Root.String()
		data, err := parse(raw, data)
		if err != nil {
			log.Println("error parsing template ", err)
			return err
		}

		log.Println(string(data))
		// if err := os.WriteFile(data.To, output, 0600); err != nil {

		// }
	}
	return nil
}

// generateTemplate - execute template with data and write to bytes buffer
func generateTemplate(t *template.Template, data Data) ([]byte, error) {
	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	if err := t.Execute(wr, &data); err != nil {
		log.Println("error generating template ", t.Name())
		return buf.Bytes(), err
	}

	if err := wr.Flush(); err != nil {
		log.Println("error flushing writer ", err)
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
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
