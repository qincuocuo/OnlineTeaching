package mbase

type mapping struct {
}

var Mapping mapping

// StrKeyInMap 判断字符串是否是map中的key
func (mapping) StrKeyInMap(key string, src map[string]string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}
