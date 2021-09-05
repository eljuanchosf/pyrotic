package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "a new command",
		Long:  "a new command",
		Run:   versionFunc,
	}
}

func versionFunc(cmd *cobra.Command, args []string) {

	log.Println(version)
}
