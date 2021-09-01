package engine

import (
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

func New(dirPath string, filePrefix string) (Core, error) {
	tmp, err := withTemplates(template.New("root"), dirPath, filePrefix)
	if err != nil {
		return Core{}, err
	}
	return Core{
		root: tmp,
	}, nil
}

// withTemplates - load templates by file path
func withTemplates(root *template.Template, fileSuffix string, dirPath string) (*template.Template, error) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return root, err
	}
	var allFiles []string
	for _, file := range files {
		filename := file.Name()
		fileLocation := fmt.Sprintf("%s/%s", dirPath, filename)
		if strings.HasSuffix(filename, fileSuffix) {
			allFiles = append(allFiles, fileLocation)
		}
	}
	root, err = root.ParseFiles(allFiles...)
	if err != nil {
		return root, err
	}
	return root, nil
}
