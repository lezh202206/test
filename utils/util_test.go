package utils

import (
	"reflect"
	"testing"
)

func TestMatchOneOf(t *testing.T) {
	type args struct {
		patterns []string
		text     string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal test",
			args: args{
				patterns: []string{`aaa(\d+)`, `hello(\d+)`},
				text:     "hello12345",
			},
			want: []string{
				"hello12345", "12345",
			},
		},
		{
			name: "normal test",
			args: args{
				patterns: []string{`aaa(\d+)`, `bbb(\d+)`},
				text:     "hello12345",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchOneOf(tt.args.text, tt.args.patterns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchOneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
