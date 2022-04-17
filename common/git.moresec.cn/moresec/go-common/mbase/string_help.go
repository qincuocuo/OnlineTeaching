package mbase

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 比较两个slice是否相等.
func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 两个数组取交集
func Intersect(ss1 []string, ss2 []string) []string {
	var result []string
	var x = 0
	temp := make(map[string]int)
	for i := 0; i < len(ss1); i++ {
		temp[ss1[i]]++
	}
	for j := 0; j < len(ss2); j++ {
		if temp[ss2[j]] > 0 {
			result = append(result, ss2[j])
			temp[ss2[j]]--
			x++
		}
	}
	return result
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
