package mbase

import (
	"testing"
)

func TestTime(t *testing.T) {
	t.Log(CurMSecond())
	t.Log(CurSecond())
	t.Log(TimeToDate())
	t.Log(TimeToDBSufix())
	t.Log(TimeToString())

	tm, _ := StrToTime("2017-11-07 21:05:38")
	t.Log(tm.Unix())
}

func TestMD5String(t *testing.T) {
	t.Log(MD5String("3", 3))
}
