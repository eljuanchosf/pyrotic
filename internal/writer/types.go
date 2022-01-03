package writer

import "sync"

type InjectClause string

const (
	InjectBefore InjectClause = "Before"
	InjectAfter  InjectClause = "After"
)

type Write struct {
	mx sync.RWMutex
	fs fileReadWrite
}

type Inject struct {
	Matcher string
	Clause  InjectClause
}

// Validate - one clause must be met
func (i *Inject) Validate() bool {
	hasClause := (i.Clause == InjectBefore || i.Clause == InjectAfter)
	if !hasClause {
		return false
	}
	return len(i.Matcher) > 0
}

type fileWrite struct {
	DryRun bool
}

var _ fileReadWrite = (*fileWrite)(nil)

type fileLog struct{}

var _ fileReadWrite = (*fileLog)(nil)
