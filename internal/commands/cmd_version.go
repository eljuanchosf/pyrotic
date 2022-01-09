package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "App version",
		Long:  "App version",
		Run:   versionFunc,
	}
}

func versionFunc(cmd *cobra.Command, args []string) {

	log.Println(version)
}
