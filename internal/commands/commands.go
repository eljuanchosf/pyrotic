package commands

import "github.com/spf13/cobra"

var (
	templatePath   = "./_templates"
	templateSuffix = ".tmpl"
)

var (
	generateName string
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
	generate := generateCmd()
	generate.PersistentFlags().StringVarP(&generateName, "name", "n", "newGeneratedName", "name of the code generation")
	rootCmd.AddCommand(generate)
	return rootCmd.Execute()
}
