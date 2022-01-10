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

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Init pyrotic",
		Long:  "Creates required setup for pyrotic",
		Run:   initFunc,
	}
}

func initFunc(cmd *cobra.Command, args []string) {
	log.Println(chalk.Green("creating initial setup"), flagTemplatePath)
	dirPath := path.Join(flagTemplatePath, "new", fmt.Sprintf("new%s", flagTemplateSuffix))
	if err := os.MkdirAll(filepath.Dir(dirPath), 0750); err != nil {
		log.Println("error creating setup", err)
		return
	}
}
