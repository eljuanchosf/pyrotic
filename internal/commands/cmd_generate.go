package commands

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
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
		log.Println("at least 1 generator name must be provided")
		return
	}

	generator := args[0]
	dirPath := filepath.Join(flagTemplatePath, generator)
	_, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println("generator not found:", generator)
		return
	}

	sharedPath := filepath.Join(flagTemplatePath, flagSharedFolder)

	log.Println(chalk.Green("running generator:"), generator)
	e, err := engine.New(flagDryrun, dirPath, sharedPath, flagTemplateSuffix)
	if err != nil {
		log.Println("error creating engine ", err)
		return
	}

	startTime := time.Now()
	err = e.Generate(engine.Data{Name: flagGeneratorName, MetaArgs: flagMetaArgs})
	if err != nil {
		return
	}
	log.Println(chalk.Green("generated in "), time.Since(startTime))
}
