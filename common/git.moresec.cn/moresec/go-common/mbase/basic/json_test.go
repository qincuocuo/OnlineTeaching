package mbase

import (
	"encoding/json"
	"testing"
)

func TestStr_JsonValid(t *testing.T) {
	testStringMap := []struct {
		JsonData string
		Result   bool
	}{
		{
			JsonData: `{"hello":"OK","hellos":[1,2,3,4,5],"score":1.23456}`,
			Result:   true,
		},
		{
			JsonData: `{"hello":123"`,
			Result:   false,
		},
	}
	for _, value := range testStringMap {
		if Json.Valid(value.JsonData) != value.Result {
			t.Error("TestStr_JsonValid result error:", value)
		}
	}
}

func TestStr_JsonLoadBytes(t *testing.T) {
	testStruct := map[string]string{
		"hello": "123",
		"test":  "ok",
	}
	compare, _ := json.Marshal(testStruct)
	for index, data := range Json.LoadBytes(testStruct) {
		if data != compare[index] {
			t.Error("TestStr_JsonLoadBytes result error:", compare, index)
		}
	}
}
