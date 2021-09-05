package parser

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func Test_withTemplates(t *testing.T) {
	type args struct {
		root       *template.Template
		fileSuffix string
		dirPath    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return inject_after.tmpl",
			args: args{
				root:       template.New("root").Funcs(defaultFuncs),
				fileSuffix: "tmpl",
				dirPath:    "../../example/_templates/fakr",
			},
			want:    "inject_after.tmpl",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := withTemplates(tt.args.root, tt.args.fileSuffix, tt.args.dirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("withTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			list := got.Templates()
			assert.Equal(t, tt.want, list[0].Name())
		})
	}
}
