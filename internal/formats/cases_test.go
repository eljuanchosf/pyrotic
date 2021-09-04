package formats

import (
	"testing"
)

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

func TestCasePascal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return SnakeCase",
			args: args{
				str: "snake_case",
			},
			want: "SnakeCase",
		},
		{
			name: "should return CamelCase",
			args: args{
				str: "camelCase",
			},
			want: "CamelCase",
		},
		{
			name: "should return KebabCase",
			args: args{
				str: "kebab-case",
			},
			want: "KebabCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CasePascal(tt.args.str); got != tt.want {
				t.Errorf("CasePascal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return snakeCase",
			args: args{
				str: "snake_case",
			},
			want: "snakeCase",
		},
		{
			name: "should return pascalCase",
			args: args{
				str: "PascalCase",
			},
			want: "pascalCase",
		},
		{
			name: "should return kebabCase",
			args: args{
				str: "kebab-case",
			},
			want: "kebabCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseCamel(tt.args.str); got != tt.want {
				t.Errorf("CaseCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
