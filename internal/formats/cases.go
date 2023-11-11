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
var matchSymbol = regexp.MustCompile("[_-]")

var titleCase = cases.Title(language.English)

func CaseSnake(str string) string {
	tmp := toSymbolCase(str, "_")
	return strings.ToLower(tmp)
}

func CasePascal(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1} ${2}")
	tmp = titleCase.String(tmp)
	return strings.ReplaceAll(tmp, " ", "")
}

func CaseKebab(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	tmp = matchAllCap.ReplaceAllString(tmp, "${1}-${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1}-${2}")
	return strings.ToLower(tmp)
}

func CaseCamel(str string) string {
	tmp := CasePascal(str)
	return lowercaseFirst(tmp)
}

func toSymbolCase(str string, sep string) string {
	tmp := splitChunks(str)
	tmp = matchSymbol.ReplaceAllString(tmp, fmt.Sprintf("${1}%s${2}", sep))
	return strings.ToLower(tmp)
}

func splitChunks(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	return matchAllCap.ReplaceAllString(tmp, "${1}-${2}")
}

func lowercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
