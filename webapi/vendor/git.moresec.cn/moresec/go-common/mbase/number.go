package mbase

import "strconv"

// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
func Atoi(s string) int32 {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i64)
}

func Atoll(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i64
}

func Atoui(s string) uint {
	ui64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(ui64)
}

func Atoull(s string) uint64 {
	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return ui64
}

func LltoA(i int64) string {
	return strconv.FormatInt(i, 10)
}
