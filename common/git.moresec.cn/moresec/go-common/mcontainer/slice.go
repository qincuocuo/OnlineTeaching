package mcontainer

import (
	"log"
	"reflect"
	"sort"
	"time"
)

type bodyWrapper struct {
	Bodys []interface{}
	by    func(p, q *interface{}) bool
}

func (acw bodyWrapper) Len() int {
	return len(acw.Bodys)
}

func (acw bodyWrapper) Swap(i, j int) {
	acw.Bodys[i], acw.Bodys[j] = acw.Bodys[j], acw.Bodys[i]
}

func (acw bodyWrapper) Less(i, j int) bool {
	return acw.by(&acw.Bodys[i], &acw.Bodys[j])
}

func sortBodyByField(bodys []interface{}, field string, order int64) {
	sort.Sort(bodyWrapper{bodys, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName(field)
		v = reflect.ValueOf(*q)
		j := v.FieldByName(field)
		switch i.Interface().(type) {
		case int:
			if order == 1 {
				return i.Int() < j.Int()
			}
			return i.Int() > j.Int()
		case int32:
			if order == 1 {
				return i.Int() < j.Int()
			}
			return i.Int() > j.Int()
		case int64:
			if order == 1 {
				return i.Int() < j.Int()
			}
			return i.Int() > j.Int()
		case string:
			if order == 1 {
				return i.String() < j.String()
			}
			return i.String() > j.String()
		case time.Time:
			if order == 1 {
				return i.Interface().(time.Time).Unix() < j.Interface().(time.Time).Unix()
			}
			return i.Interface().(time.Time).Unix() > j.Interface().(time.Time).Unix()
		default:
			log.Printf("this type cant't be sorted\n")
			return true
		}
	}})
}

// Sorter is ...
type Sorter struct {
	Field string
	Aesc  int64
}

// StructSliceQuery is use like sql to filter sort paging data
func StructSliceQuery(data []interface{}, where map[string]interface{}, sort Sorter, limit, skip int64) []interface{} {
	// filter
	rawResult := data
	for k, w := range where {
		var tempResult []interface{}
		func(temp []interface{}, key string, value interface{}) {
			for _, v := range rawResult {
				rv := reflect.ValueOf(v)
				if rv.FieldByName(k).Interface() == w {
					tempResult = append(tempResult, v)
				}
			}
			rawResult = tempResult
		}(rawResult, k, w)
	}

	// sort
	sortBodyByField(rawResult, sort.Field, sort.Aesc)

	// paging
	end := skip + limit
	if end > int64(len(rawResult)) {
		end = int64(len(rawResult))
	}
	return rawResult[skip:end]
}
