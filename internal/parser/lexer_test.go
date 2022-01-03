package parser

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hydrateData(t *testing.T) {
	type args struct {
		meta []string
		data TemplateData
	}
	tests := []struct {
		name string
		args args
		want TemplateData
	}{
		{
			name: "should return inject before",
			args: args{
				meta: []string{
					"inject: true",
					"before: // deepak",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: false,
				Inject: true,
				Before: "// deepak",
				After:  "",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should return inject after",
			args: args{
				meta: []string{
					"inject: true",
					"after: // deepak",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: false,
				Inject: true,
				Before: "",
				After:  "// deepak",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should return append",
			args: args{
				meta: []string{
					"append: true",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: true,
				Inject: false,
				Before: "",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should return to ",
			args: args{
				meta: []string{
					"to: example/screen/foo",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "example/screen/foo",
				Append: false,
				Inject: false,
				Before: "",
				After:  "",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should return to ",
			args: args{
				meta: []string{
					"block: steel",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: false,
				Inject: false,
				Before: "",
				After:  "",
				Output: nil,
				Meta: map[string]string{
					"block": "steel",
				},
			},
		},
		{
			name: "should return to  and remove white spaces",
			args: args{
				meta: []string{
					"  to  : steel  ",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "steel",
				Append: false,
				Inject: false,
				Before: "",
				After:  "",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should skip parse and return default",
			args: args{
				meta: []string{
					"  to  steel  ",
				},
				data: TemplateData{},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: false,
				Inject: false,
				Before: "",
				After:  "",
				Output: nil,
				Meta:   map[string]string{},
			},
		},
		{
			name: "should skip parse and return with o",
			args: args{
				meta: []string{
					"  to  steel  ",
				},
				data: TemplateData{
					Meta: map[string]string{
						"foo": "bar",
					},
				},
			},
			want: TemplateData{
				Name:   "",
				To:     "",
				Append: false,
				Inject: false,
				Before: "",
				After:  "",
				Output: nil,
				Meta: map[string]string{
					"foo": "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hydrateData(tt.args.meta, tt.args.data)
			gotJSON, err := json.Marshal(&got)
			assert.NoError(t, err)
			wantJSON, err := json.Marshal(tt.want)
			assert.NoError(t, err)
			assert.JSONEq(t, string(wantJSON), string(gotJSON))
		})
	}
}

func Test_extractMeta(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name   string
		args   args
		meta   []string
		output string
	}{
		{
			name: "should return meta block",
			args: args{
				output: `---
				to: foo
				---
				`,
			},
			meta:   []string{"to: foo"},
			output: "",
		},
		{
			name: "should empty if no block",
			args: args{
				output: `
				to: foo
				`,
			},
			meta:   []string{},
			output: "",
		},
		{
			name: "should return meta and block",
			args: args{
				output: `---
				append: true
				---
				blah
				`,
			},
			meta:   []string{"append: true"},
			output: "blah",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := extractMeta(tt.args.output)
			if !reflect.DeepEqual(got, tt.meta) {
				t.Errorf("extractMeta() got = %v, want %v", got, tt.meta)
			}
			if strings.TrimSpace(got1) != tt.output {
				t.Errorf("extractMeta() got1 = %v, want %v", got1, tt.output)
			}
		})
	}
}
