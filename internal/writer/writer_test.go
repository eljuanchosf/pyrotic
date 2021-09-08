package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_injectIntoData(t *testing.T) {
	type args struct {
		name   string
		source []byte
		data   []byte
		inject Inject
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "inject before token",
			args: args{
				name:   "",
				source: []byte("fall of  // token"),
				data:   []byte("fart"),
				inject: Inject{
					Before: "// token",
					After:  "",
				},
			},
			want: []byte("fall of fart\n// token"),
		},
		{
			name: "inject after token",
			args: args{
				name:   "",
				source: []byte("fall of // token"),
				data:   []byte("fart"),
				inject: Inject{
					Before: "",
					After:  "// token",
				},
			},
			want: []byte("fall of // tokenfart"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeOutputs(tt.args.name, tt.args.source, tt.args.data, tt.args.inject)
			assert.Equal(t, string(tt.want), string(got))
		})
	}
}
