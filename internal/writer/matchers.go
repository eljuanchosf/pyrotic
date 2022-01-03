package writer

import (
	"fmt"
	"strings"
)

func mergeInjection(source, dataInjection []byte, inject Inject) ([]byte, error) {
	var splitByMatcher []string
	if err := inject.Validate(); err != nil {
		return source, err
	}

	switch inject.Clause {
	case InjectAfter:
		splitByMatcher = strings.SplitAfter(string(source), inject.Matcher)
		if len(splitByMatcher) != 2 {
			return source, ErrNoMatchingExpression
		}
	case InjectBefore:
		idx := strings.Index(string(source), inject.Matcher)
		if idx == -1 {
			return source, ErrNoMatchingExpression
		}
		splitByMatcher = []string{
			string(source[:(idx - 1)]),
			fmt.Sprintf("\n%s", string(source[idx:])),
		}
	}

	formatedOutput := strings.Join([]string{
		splitByMatcher[0],
		string(dataInjection),
		splitByMatcher[1],
	}, "")
	return []byte(formatedOutput), nil
}
