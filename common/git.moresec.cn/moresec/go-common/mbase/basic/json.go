package mbase

import "encoding/json"

type mjson struct {}
var Json mjson

// Valid  校验Json字符串是否合法
func (mjson) Valid(jsonStr string) (valid bool) {
	input := []byte(jsonStr)
	var x struct{}
	err := json.Unmarshal(input, &x)
	if err != nil {
		valid = false
	} else {
		valid = true
	}
	return
}

// LoadBytes  字节转Json字符串
func (mjson) LoadBytes(src interface{}) (jsonStr []byte) {
	jsonStr, _ = json.Marshal(src)
	return
}