package commands

import (
	"github.com/spf13/cobra"
)

var (
	templatePath   = "_templates"
	templateSuffix = ".tmpl"
	dryrun         = false
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
	rootCmd.PersistentFlags().StringVarP(&templatePath, "path", "p", "_templates", "templates path eg: _templates")
	rootCmd.PersistentFlags().StringVarP(&templateSuffix, "extension", "x", ".tmpl", "template extension eg: *.tmpl")
	rootCmd.PersistentFlags().BoolVarP(&dryrun, "dry-run", "d", false, "run the generator in dry run mode")
}

func Execute() error {
	generate := generateCmd()
	generate.PersistentFlags().StringVarP(&generateName, "name", "n", "newGeneratedName", "name of the code generation")
	rootCmd.AddCommand(generate)
	// add commands
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(newCmd())
	return rootCmd.Execute()
}
