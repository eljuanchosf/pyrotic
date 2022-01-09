package commands

import "errors"

var (
	ErrNoArguments       = errors.New("no arguments provided")
	ErrGeneratorNotFound = errors.New("generator not found")
	Err                  = errors.New("err")
)
