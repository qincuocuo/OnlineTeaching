package mbase

import (
	"testing"
)

func TestStr_Base64Decrypt(t *testing.T) {
	stringTestTable := map[string]string{
		"5rWL6K+V6YCa6L+H=": "测试通过",
		"5rWL6K+V6YCa6L+H":  "测试通过",
	}
	for key, value := range stringTestTable {
		if String.Base64Decrypt(key) != value {
			t.Error("TestStr_Base64Decrypt result error", key, value)
		}
	}
}
func TestStr_Compare(t *testing.T) {
	TestTable := []struct {
		StrA   string
		StrB   string
		Result bool
	}{
		{
			StrA:   "Hello",
			StrB:   "hello",
			Result: false,
		},
		{
			StrA:   "Hello",
			StrB:   "Hello",
			Result: true,
		},
	}
	for _, test := range TestTable {
		if String.Compare(test.StrA, test.StrB) != test.Result {
			t.Error("TestStr_Compare result error:", test)
		}
	}
}
func TestStr_Md5(t *testing.T) {
	stringTest := "I want a md5 string .*?!@#"
	actual := String.Md5(stringTest)
	if actual == "" {
		t.Error("TestStr_Md5 result error:", stringTest, actual)
	}
}
func TestStr_ToInt(t *testing.T) {
	stringTestMap := map[string]int{"1": 1, "0": 0, "9523": 9523, "3389": 3389, "123123123": 123123123}
	for key, value := range stringTestMap {
		if String.ToInt(key) != value {
			t.Error("TestStr_ToInt result error:", key, value)
		}
	}
}
func TestStr_Int32ToString(t *testing.T) {
	stringTestMap := map[int32]string{1: "1", 0: "0", 9523: "9523", 3389: "3389", 123123123: "123123123"}
	for key, value := range stringTestMap {
		if String.Int32ToString(key) != value {
			t.Error("TestStr_Int32ToString result error:", key, value)
		}
	}
}

func TestStr_GetPasswordStrength(t *testing.T) {
	//四个等级: 0 <1 <2 <3
	//四个强度元素: 大写字母 小写字母 数字 特殊字符
	// 0 - 长度小于8 || 强度元素小于等于1个
	// 1 - 强度元素小于等于2个
	// 2 - 强度元素小于等于3个
	// 3 - 强度元素小于等于4个
	testStringMap := map[string]int{
		"qwer":      0,
		"qWer":      0,
		"qW1r":      0,
		"qW1@":      0,
		"qWer123":   0,
		"qwer1234":  1,
		"qwer@1234": 2,
		"qWer@1234": 3,
	}
	for key, value := range testStringMap {
		if String.GetPasswordStrength(key) != value {
			t.Error("TestStr_GetPasswordStrength result error:", key, value)
		}
	}
}
