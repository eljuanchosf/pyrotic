package formats

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchSymbol = regexp.MustCompile(`[\s_-]`)

var titleCase = cases.Title(language.English)

func CaseSnake(str string) string {
	tmp := replaceStringWithSep(str, "_")
	return strings.ToLower(tmp)
}

func CasePascal(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1} ${2}")
	tmp = titleCase.String(tmp)
	return strings.ReplaceAll(tmp, " ", "")
}

func CaseKebab(str string) string {
	tmp := replaceStringWithSep(str, "-")
	return strings.ToLower(tmp)
}

func CaseCamel(str string) string {
	tmp := CasePascal(str)
	return lowercaseFirst(tmp)
}

func replaceStringWithSep(str, sep string) string {
	expression := fmt.Sprintf("${1}%s${2}", sep)
	if matchSymbol.Match([]byte(str)) {
		return matchSymbol.ReplaceAllString(str, expression)
	}
	tmp := matchFirstCap.ReplaceAllString(str, expression)
	fmt.Println("print", tmp)

	return tmp
}

func lowercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
