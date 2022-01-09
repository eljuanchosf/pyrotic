package commands

import (
	"github.com/spf13/cobra"
)

var (
	// root vars
	flagTemplatePath   = "_templates"
	flagTemplateSuffix = ".tmpl"
	flagMetaArgs       = ""
	flagDryrun         = false
	flagGeneratorName  string
	flagSharedFolder   = "shared"
)

const (
	version = "v.dev-1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "pyrotic [COMMAND] --[FLAGS]",
	Short: "Simple code generation",
	Long:  `Simple code generation`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flagTemplatePath, "path", "p", "_templates", "Templates directory path, defaults to _templates at the root of the project.")
	rootCmd.PersistentFlags().StringVarP(&flagTemplateSuffix, "extension", "x", ".tmpl", "Template file extension, defaults to .tmpl")
	rootCmd.PersistentFlags().StringVarP(&flagMetaArgs, "meta", "m", "", "Pass meta arguments to template. Meta arguments passed via command line will overwrite emplate args. Use a comma delimiter to provide multiple arguments, eg: --meta foo=bar,bin=baz")
	rootCmd.PersistentFlags().StringVarP(&flagSharedFolder, "shared", "s", "Shared", "shared template folder name, defaults to shared eg: shared")
	rootCmd.PersistentFlags().BoolVarP(&flagDryrun, "dry-run", "d", false, "Run the generator in dry run mode")
}

func Execute() error {
	generate := generateCmd()
	generate.PersistentFlags().StringVarP(&flagGeneratorName, "name", "n", "newGeneratedName", "Name of the newly generated code eg: --name new-fakr")
	rootCmd.AddCommand(generate)
	// add commands
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(newCmd())
	return rootCmd.Execute()
}
