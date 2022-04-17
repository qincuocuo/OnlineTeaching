package mbase

import (
	"reflect"
	"testing"
)

func TestStringSliceEqualBCE(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{[]string{"a", "b", "c"}, []string{"a", "b", "c"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceEqualBCE(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StringSliceEqualBCE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		ss1 []string
		ss2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{[]string{"a", "b", "c"}, []string{"a"}}, want: []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.ss1, tt.args.ss2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
