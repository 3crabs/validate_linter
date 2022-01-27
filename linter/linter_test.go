package linter

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, "/Users/vladimir/go/src/validate-linter/testdata/test_string", NewAnalyzer())
}

func Test_getValidateRule(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get gt=0",
			args: args{tag: "`validate:\"gt=0\"`"},
			want: "gt=0",
		},
		{
			name: "get gte=0",
			args: args{tag: "`validate:\"gte=0\"`"},
			want: "gte=0",
		},
		{
			name: "get omitempty",
			args: args{tag: "`validate:\"omitempty\"`"},
			want: "omitempty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValidateRule(tt.args.tag); got != tt.want {
				t.Errorf("getValidateRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkString(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "string field gt=0",
			args:    args{tag: "`validate:\"gt=0\"`"},
			wantErr: false,
		},
		{
			name:    "string field omitempty",
			args:    args{tag: "`validate:\"omitempty\"`"},
			wantErr: false,
		},
		{
			name:    "string field gte=0 -> omitempty",
			args:    args{tag: "`validate:\"gte=0\"`"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkString(tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("checkString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
