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
	Short: "simple code generation",
	Long:  `simple code generation`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flagTemplatePath, "path", "p", "_templates", "templates path eg: _templates")
	rootCmd.PersistentFlags().StringVarP(&flagTemplateSuffix, "extension", "x", ".tmpl", "template extension eg: *.tmpl")
	rootCmd.PersistentFlags().StringVarP(&flagMetaArgs, "meta", "m", "", "pass meta arguments to template. Meta arguments passed via command line will overwrite emplate args. Use a comma delimiter to provide multiple arguments, eg: --meta foo=bar,bin=baz")
	rootCmd.PersistentFlags().StringVarP(&flagSharedFolder, "shared", "s", "shared", "shared template folder name, defaults to shared eg: _templates/shared")
	rootCmd.PersistentFlags().BoolVarP(&flagDryrun, "dry-run", "d", false, "run the generator in dry run mode")
}

func Execute() error {
	generate := generateCmd()
	generate.PersistentFlags().StringVarP(&flagGeneratorName, "name", "n", "newGeneratedName", "name of the code generator eg: --name gen")
	rootCmd.AddCommand(generate)
	// add commands
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(newCmd())
	return rootCmd.Execute()
}
