package commands

import "github.com/spf13/cobra"

var (
	templatePath   = "./_templates"
	templateSuffix = ".tmpl"
)

var rootCmd = &cobra.Command{
	Use:   "pyrotic [COMMAND] --[FLAGS]",
	Short: "simple code generation",
	Long:  `simple code generation`,
}

func init() {}

func Execute() error {
	return rootCmd.Execute()
}
