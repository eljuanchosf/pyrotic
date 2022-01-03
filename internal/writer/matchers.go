package writer

import (
	"fmt"
	"log"
	"strings"
)

func mergeOutputs(name string, source, output []byte, inject Inject) []byte {
	var splitByMatcher []string
	if !inject.Validate() {
		log.Printf("at least 1 injection clause must not be empty, matcher: [%s], clause: [%s]", inject.Matcher, inject.Clause)
		return source
	}
	switch inject.Clause == InjectAfter {
	case true:
		splitByMatcher = strings.SplitAfter(string(source), inject.Matcher)
		if len(splitByMatcher) != 2 {
			log.Printf("injection token %s is not found in file %s", inject.Matcher, name)
			return source
		}
	default:
		idx := strings.Index(string(source), inject.Matcher)
		if idx == -1 {
			log.Printf("injection token %s is not found in file %s", inject.Matcher, name)
			return source
		}
		splitByMatcher = []string{
			string(source[:(idx - 1)]),
			fmt.Sprintf("\n%s", string(source[idx:])),
		}
	}

	formatedOutput := strings.Join([]string{
		splitByMatcher[0],
		string(output),
		splitByMatcher[1],
	}, "")
	return []byte(formatedOutput)
}
