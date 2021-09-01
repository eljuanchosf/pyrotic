package commands

import (
	"log"

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
		log.Println("at least 1 template name must be provided")
		return
	}

}
