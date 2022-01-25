package formats

import "strings"

func SplitByDelimiter(value string, delimiter string) []string {
	return strings.Split(value, delimiter)
}

func SplitAfterDelimiter(value string, delimiter string) []string {
	return strings.SplitAfter(value, delimiter)
}

func Contains(value string, subStr string, n int) bool {
	return strings.Contains(value, subStr)
}

func HasPrefix(value string, prefix string, n int) bool {
	return strings.HasPrefix(value, prefix)
}

func HasSuffix(value string, suffix string, n int) bool {
	return strings.HasSuffix(value, suffix)
}
