package mcontainer

import (
	"reflect"
	"testing"
	"time"
)

type TestUser struct {
	Name     string
	Age      int64
	Birthday time.Time
}

func getTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}

func TestStructSliceQuery(t *testing.T) {
	type args struct {
		data  []interface{}
		where map[string]interface{}
		sort  Sorter
		limit int64
		skip  int64
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "case1",
			args: args{
				data:  []interface{}{TestUser{"cc", 2, getTime("2019-08-03 15:04:05")}, TestUser{"cc", 3, getTime("2019-08-08 15:04:05")}, TestUser{"mm", 5, getTime("2019-08-02 15:04:05")}, TestUser{"cc", 1, getTime("2019-08-01 15:04:05")}},
				where: map[string]interface{}{"Name": "cc"},
				sort:  Sorter{"Birthday", -1},
				limit: 10,
				skip:  0,
			},
			want: []interface{}{TestUser{"cc", 3, getTime("2019-08-08 15:04:05")}, TestUser{"cc", 2, getTime("2019-08-03 15:04:05")}, TestUser{"cc", 1, getTime("2019-08-01 15:04:05")}},
		},
		{
			name: "case2",
			args: args{
				data:  []interface{}{TestUser{"cc", 2, getTime("2019-08-03 15:04:05")}, TestUser{"cc", 3, getTime("2019-08-08 15:04:05")}, TestUser{"mm", 5, getTime("2019-08-02 15:04:05")}, TestUser{"cc", 1, getTime("2019-08-01 15:04:05")}},
				where: map[string]interface{}{"Name": "cc"},
				sort:  Sorter{"Age", 1},
				limit: 10,
				skip:  0,
			},
			want: []interface{}{TestUser{"cc", 1, getTime("2019-08-01 15:04:05")}, TestUser{"cc", 2, getTime("2019-08-03 15:04:05")}, TestUser{"cc", 3, getTime("2019-08-08 15:04:05")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructSliceQuery(tt.args.data, tt.args.where, tt.args.sort, tt.args.limit, tt.args.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructSliceQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
