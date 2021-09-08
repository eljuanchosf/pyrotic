package writer

import "sync"

type Write struct {
	mx sync.RWMutex
	fs fileReadWrite
}

type Inject struct {
	Before string
	After  string
}

type fileWrite struct {
	DryRun bool
}

var _ fileReadWrite = (*fileWrite)(nil)

type fileLog struct{}

var _ fileReadWrite = (*fileLog)(nil)
