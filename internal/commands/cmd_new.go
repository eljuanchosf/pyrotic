package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/spf13/cobra"
)

func newCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "create generator",
		Long:  "create a new generator",
		Run:   new,
	}
}

func new(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("at least 1 generator name must be provided")
		return
	}
	generator := args[0]
	log.Println(chalk.Green("creating new generator"), templatePath)
	dirPath := path.Join(templatePath, generator, fmt.Sprintf("%s%s", generator, templateSuffix))
	if err := os.MkdirAll(filepath.Dir(dirPath), 0750); err != nil {
		log.Println("error creating", err)
		return
	}

	file, err := os.Create(dirPath)
	if err != nil {
		log.Println("error creating base template ", err)
		return
	}
	defer file.Close()
	if _, err := file.Write([]byte(newGeneratorTemplate)); err != nil {
		log.Println("error creating file")
		return
	}
}

var newGeneratorTemplate = `---
to:
---
package main

func main() {}
`
