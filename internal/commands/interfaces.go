package commands

import "github.com/code-gorilla-au/pyrotic/internal/engine"

type Generator interface {
	Generate(data engine.Data) error
}
