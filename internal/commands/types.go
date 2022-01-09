package commands

import "github.com/spf13/cobra"

type cmdErrFunc = func(cmd *cobra.Command, args []string) error
