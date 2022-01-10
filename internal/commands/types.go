package commands

import "github.com/spf13/cobra"

type cmdFunc = func(cmd *cobra.Command, args []string)
