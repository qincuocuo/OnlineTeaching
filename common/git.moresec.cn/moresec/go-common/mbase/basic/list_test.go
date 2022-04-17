package mbase

import (
	"github.com/globalsign/mgo/bson"
	"reflect"
	"testing"
)

func TestRemoveDuplicate_StringList(t *testing.T) {
	testStringList := []string{"a", "b", "a", "b"}
	resultStringList := []string{"a", "b"}
	if !reflect.DeepEqual(resultStringList, RemoveDuplicate.StringList(testStringList)) {
		t.Error("TestRemoveDuplicate_StringList result error", testStringList, resultStringList)
	}
}

func TestRemoveDuplicate_IntList(t *testing.T) {
	testIntList := []int{1, 2, 2, 1, 1}
	resultIntList := []int{1, 2}
	if !reflect.DeepEqual(resultIntList, RemoveDuplicate.IntList(testIntList)) {
		t.Error("TestRemoveDuplicate_IntList result error", testIntList, resultIntList)
	}
}

func TestRemoveDuplicate_Int32List(t *testing.T) {
	testInt32List := []int32{1, 2, 2, 1, 1}
	resultInt32List := []int32{1, 2}
	if !reflect.DeepEqual(resultInt32List, RemoveDuplicate.Int32List(testInt32List)) {
		t.Error("TestRemoveDuplicate_Int32List result error", testInt32List, resultInt32List)
	}
}

func TestRemoveDuplicate_Int64List(t *testing.T) {
	testInt64List := []int64{1, 2, 2, 1, 1}
	resultInt64List := []int64{1, 2}
	if !reflect.DeepEqual(resultInt64List, RemoveDuplicate.Int64List(testInt64List)) {
		t.Error("TestRemoveDuplicate_Int32List result error", testInt64List, resultInt64List)
	}
}
func TestRemoveDuplicate_Float32List(t *testing.T) {
	testFloat32List := []float32{1.1, 2.2, 2.2, 1.1, 1.1}
	resultFloat32List := []float32{1.1, 2.2}
	if !reflect.DeepEqual(resultFloat32List, RemoveDuplicate.Float32List(testFloat32List)) {
		t.Error("TestRemoveDuplicate_Int32List result error", testFloat32List, resultFloat32List)
	}
}
func TestRemoveDuplicate_Float64List(t *testing.T) {
	testFloat64List := []float64{1.1, 2.2, 2.2, 1.1, 1.1}
	resultFloat64List := []float64{1.1, 2.2}
	if !reflect.DeepEqual(resultFloat64List, RemoveDuplicate.Float64List(testFloat64List)) {
		t.Error("TestRemoveDuplicate_Int32List result error", testFloat64List, resultFloat64List)
	}
}

func TestListType_Int32ToInt(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5, 6}
	result := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, ListType.Int32ToInt(input)) {
		t.Error("TestListType_Int32ToInt result error", input, result)
	}
}

func TestListType_IntToInt32(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := []int32{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, ListType.IntToInt32(input)) {
		t.Error("TestListType_Int32ToInt result error", input, result)
	}
}

func TestListType_StringToObj(t *testing.T) {
	var input []string
	var result []bson.ObjectId
	for i := 0; i < 3; i++ {
		id := bson.NewObjectId()
		input = append(input, id.Hex())
		result = append(result, id)

	}
	if !reflect.DeepEqual(result, ListType.StringToObj(input)) {
		t.Error("TestListType_StringToObj result error", result, input)
	}
}

func TestInList_IntList(t *testing.T) {
	testStruct := []struct {
		Key    int
		List   []int
		Result bool
	}{
		{
			Key:    1,
			List:   []int{1, 2, 3, 4, 5, 6},
			Result: true,
		},
		{
			Key:    0,
			List:   []int{1, 2, 3, 4, 5, 6},
			Result: false,
		},
		{
			Key:    1,
			List:   []int{},
			Result: false,
		},
	}
	for _, item := range testStruct {
		if item.Result != InList.IntList(item.Key, item.List) {
			t.Error("TestInList_IntList result error", item)
		}
	}
}
