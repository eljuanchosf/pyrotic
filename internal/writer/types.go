package writer

import "sync"

type Write struct {
	mx sync.RWMutex
	fs fileReadWrite
}

type fileWrite struct {
	DryRun bool
}

var _ fileReadWrite = (*fileWrite)(nil)

type fileLog struct{}

var _ fileReadWrite = (*fileLog)(nil)

type inject struct {
	After   bool
	Matcher string
}
