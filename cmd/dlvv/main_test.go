package main

import (
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		args      []string
		wantArgs  []string
		wantExtra []string
	}{
		{
			[]string{"-v"},
			[]string{},
			[]string{"-test.v"},
		},
		{
			[]string{"-run", "Foo"},
			[]string{},
			[]string{"-test.run", "Foo"},
		},
		{
			[]string{"-run=Foo"},
			[]string{},
			[]string{"-test.run=Foo"},
		},
		{
			[]string{"-count", "10"},
			[]string{},
			[]string{"-test.count", "10"},
		},
		{
			[]string{"-count=10"},
			[]string{},
			[]string{"-test.count=10"},
		},
	}

	for _, tt := range tests {
		gotArgs, gotExtra := splitTestArgs(tt.args)
		if !slices.Equal(tt.wantArgs, gotArgs) {
			t.Errorf("wants args %v but got %v", tt.wantArgs, gotArgs)
		}
		if !slices.Equal(tt.wantExtra, gotExtra) {
			t.Errorf("wants extra %v but got %v", tt.wantExtra, gotExtra)
		}
	}
}
