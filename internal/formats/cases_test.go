package formats

import "testing"

func TestCaseSnake(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "camel should return snake_case",
			args: args{
				str: "snakeCase",
			},
			want: "snake_case",
		},
		{
			name: "pascal should return pascal_case",
			args: args{
				str: "PascalCase",
			},
			want: "pascal_case",
		},
		{
			name: "kebab should return kebab_case",
			args: args{
				str: "kebab-case",
			},
			want: "kebab_case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseSnake(tt.args.str); got != tt.want {
				t.Errorf("CaseSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
