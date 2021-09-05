package parser

import (
	"encoding/json"
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
				Meta:   map[string]interface{}{},
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
				Meta:   map[string]interface{}{},
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
				Meta:   map[string]interface{}{},
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
				Meta: map[string]interface{}{
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
				Meta:   map[string]interface{}{},
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
				Meta:   map[string]interface{}{},
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
