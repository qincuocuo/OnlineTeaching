package mbase

import "reflect"

func InSliceInt(v int, array []int) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceInt8(v int8, array []int8) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceInt16(v int16, array []int16) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceInt32(v int32, array []int32) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceInt64(v int64, array []int64) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceUint(v uint, array []uint) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceUint8(v uint8, array []uint8) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceUint16(v uint16, array []uint16) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceUint32(v uint32, array []uint32) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceUint64(v uint64, array []uint64) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceFloat32(v float32, array []float32) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceFloat64(v float64, array []float64) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceString(v string, array []string) bool {
	for _, item := range array {
		if v == item {
			return true
		}
	}
	return false
}

func InSliceAny(val interface{}, array interface{}) bool {
	if array == nil {
		return false
	}

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		if s.IsNil() {
			return false
		}

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
