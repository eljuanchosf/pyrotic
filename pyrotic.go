package main

import (
	"os"

	"github.com/code-gorilla-au/pyrotic/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
