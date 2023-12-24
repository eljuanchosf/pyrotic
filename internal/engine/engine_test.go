package engine

import (
	"testing"

	"github.com/code-gorilla-au/odize"
)

func Test_generateMeta_should_return_empty(t *testing.T) {

	got := generateMeta("")
	odize.AssertEqual(t, map[string]string{}, got)
}

func Test_generateMeta_should_should_meta_map(t *testing.T) {

	got := generateMeta("foo=bar,bin=baz")
	odize.AssertEqual(t, map[string]string{
		"foo": "bar",
		"bin": "baz",
	}, got)
}

func Test_generateMeta_should_should_meta_map_with_no_spaces(t *testing.T) {

	got := generateMeta("foo = bar , bin = baz")
	odize.AssertEqual(t, map[string]string{
		"foo": "bar",
		"bin": "baz",
	}, got)
}
