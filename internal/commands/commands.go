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

func init() {
	rootCmd.PersistentFlags().StringVarP(&templatePath, "path", "p", "./_templates", "_templates path")
}

func Execute() error {
	rootCmd.AddCommand(generateCmd())
	return rootCmd.Execute()
}
