package parser

import (
	"strings"
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
		want    int
		wantErr bool
	}{
		{
			name: "should return inject_after.tmpl",
			args: args{
				root:       template.New("root").Funcs(defaultFuncs),
				fileSuffix: "tmpl",
				dirPath:    "../../example/_templates/fakr",
			},
			want:    5,
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
			assert.Equal(t, tt.want, len(list))
		})
	}
}

func TestTmplEngine_Parse_should_render(t *testing.T) {
	strTmp := `---
	to: elo
	---
	blah
	`
	expected := TemplateData{
		Name: "hello",
	}
	root := template.New("root")
	root, err := root.Parse(strTmp)
	assert.NoError(t, err)
	te := &TmplEngine{
		root:  root,
		funcs: defaultFuncs,
	}
	data, err := te.Parse(expected)
	assert.NoError(t, err)
	assert.Equal(t, expected.Name, data[0].Name)
	assert.Equal(t, "elo", data[0].To)
	assert.Equal(t, "blah", strings.TrimSpace(string(data[0].Output)))
}
func TestTmplEngine_Parse_missing_funcs_should_fail_on_meta_parse(t *testing.T) {
	strTmp := `---
	to: {{ "elo" | caseSnake }}
	---
	blah
	`
	expected := TemplateData{
		Name: "hello",
	}
	root := template.New("root").Funcs(defaultFuncs)
	root, err := root.Parse(strTmp)
	assert.NoError(t, err)
	te := &TmplEngine{
		root: root,
	}
	_, err = te.Parse(expected)
	assert.Error(t, err)
}

func TestTmplEngine_Parse_missing_funcs_should_fail_on_template_parse(t *testing.T) {
	strTmp := `---
	to: elo
	---
	blah {{ "foo" | caseSnake }}
	`
	expected := TemplateData{
		Name: "hello",
	}
	root := template.New("root").Funcs(defaultFuncs)
	root, err := root.Parse(strTmp)
	assert.NoError(t, err)
	te := &TmplEngine{
		root: root,
	}
	_, err = te.Parse(expected)
	assert.Error(t, err)
}
