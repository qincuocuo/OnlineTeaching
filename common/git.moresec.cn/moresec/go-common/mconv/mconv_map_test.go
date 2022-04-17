package mconv

import (
	"testing"
)

func TestMap(t *testing.T) {
	test1 := map[string]string{
		"a": "1",
		"b": "2",
	}
	result1 := Map(test1)
	if len(result1) != len(test1) {
		t.Errorf("Map(%v) != %v", test1, result1)
	}

	type Test2Type struct {
		Type int `json:"type"`
	}
	test2 := struct {
		Name string    `json:"name"`
		Age  string    `json:"age"`
		Data Test2Type `json:"data"`
	}{
		Name: "test",
		Age:  "12",
		Data: Test2Type{
			Type: 1,
		},
	}
	result2 := Map(test2, "data")
	x, ok := result2["data"].(Test2Type)
	if !ok {
		t.Errorf("Map(%v) != %v", test2, result2)
	}
	if x.Type != 1 {
		t.Errorf("Map(%v) != %v", test2, result2)
	}
}
