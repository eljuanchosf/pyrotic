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

// Validate - exactly 1 clause must be met. Matcher must not be empty
func (i *Inject) Validate() error {
	hasClause := (i.Clause == InjectBefore || i.Clause == InjectAfter)
	if !hasClause {
		return ErrNoMatchingClause
	}
	if len(i.Matcher) <= 0 {
		return ErrNoMatchingExpression
	}
	return nil
}

type fileWrite struct {
	DryRun bool
}

var _ fileReadWrite = (*fileWrite)(nil)

type fileLog struct{}

var _ fileReadWrite = (*fileLog)(nil)
