package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/code-gorilla-au/pyrotic/internal/engine"
	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "generate template",
		Long:  "generate tempate by argument",
		Run:   generate,
	}
}

func generate(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("at least 1 template name must be provided")
		return
	}
	generator := args[0]
	dirPath := filepath.Join(templatePath, generator)
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println("generator not found ", err)
		return
	}
	log.Println(files[0].Name())
	log.Println(dirPath)
	e, err := engine.New(dirPath, templateSuffix)
	if err != nil {
		log.Println("error creating engine ", err)
		return
	}

	output, err := e.Generate(engine.Data{Name: "goo"})
	if err != nil {
		log.Println("error generating templates ", err)
		return
	}

	log.Println(string(output))

}
