package mbase

import "testing"

func TestInSliceAny(t *testing.T) {
	type testStruct struct {
		Name string
	}

	type stringAlias string

	type args struct {
		val   interface{}
		array interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				val:   "a",
				array: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				val:   "d",
				array: []string{"a", "b", "c"},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				val:   10086,
				array: []string{"a", "b", "c"},
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				val:   "a",
				array: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				val:   "a",
				array: nil,
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				val:   "a",
				array: ([]string)(nil),
			},
			want: false,
		},
		{
			name: "test7",
			args: args{
				val: testStruct{
					Name: "a",
				},
				array: []testStruct{
					{
						Name: "a",
					},
				},
			},
			want: true,
		},
		{
			name: "test8",
			args: args{
				val:   "foo",
				array: []stringAlias{stringAlias("foo")},
			},
			want: false,
		},
		{
			name: "test9",
			args: args{
				val:   stringAlias("foo"),
				array: []stringAlias{stringAlias("foo")},
			},
			want: true,
		},
		{
			name: "test10",
			args: args{
				val:   nil,
				array: []interface{}{nil},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSliceAny(tt.args.val, tt.args.array); got != tt.want {
				t.Errorf("InSliceAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
