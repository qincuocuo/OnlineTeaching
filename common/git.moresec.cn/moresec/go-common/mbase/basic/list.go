package mbase

import (
	"github.com/globalsign/mgo/bson"
)

type removeDuplicate struct{}

var RemoveDuplicate removeDuplicate

// StringList 字符串列表去重
func (removeDuplicate) StringList(src []string) (stringList []string) {
	temp := map[string]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

// IntList 整型列表去重
func (removeDuplicate) IntList(src []int) (stringList []int) {
	temp := map[int]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

// Int32List 32位整型列表去重
func (removeDuplicate) Int32List(src []int32) (stringList []int32) {
	temp := map[int32]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

// Int64List 64位整型列表去重
func (removeDuplicate) Int64List(src []int64) (stringList []int64) {
	temp := map[int64]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

// Float64List 64位浮点型列表去重
func (removeDuplicate) Float64List(src []float64) (stringList []float64) {
	temp := map[float64]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

// Float32List 32位浮点型列表去重
func (removeDuplicate) Float32List(src []float32) (stringList []float32) {
	temp := map[float32]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			stringList = append(stringList, item)
		}
	}
	return
}

type listType struct{}

var ListType listType

// IntToInt32 整型列表转化为32位整型列表
func (listType) IntToInt32(src []int) (int32List []int32) {
	for _, item := range src {
		int32List = append(int32List, int32(item))
	}
	return
}

// Int32ToInt 32位整型列表转化为整型列表
func (listType) Int32ToInt(src []int32) []int {
	intList := make([]int, 0)
	for _, item := range src {
		intList = append(intList, int(item))
	}
	return intList
}

// StringToObj 字符串列表转化为bson对象ID列表
func (listType) StringToObj(src []string) (objList []bson.ObjectId) {
	for _, item := range src {
		objList = append(objList, bson.ObjectIdHex(item))
	}
	return
}

type inList struct{}

var InList inList

// IntList 判断整型数据是否在整型列表中
func (inList) IntList(key int, array []int) bool {
	for _, item := range array {
		if key == item {
			return true
		}
	}
	return false
}
// StringList 判断字符串数据是否在字符串列表中
func (inList) StringList(key string, array []string) bool {
	for _, item := range array {
		if key == item {
			return true
		}
	}
	return false
}
