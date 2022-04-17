package mbase

import (
	"testing"
	"time"
)

func TestTimeToStr(t *testing.T) {
	timeTest := time.Now()
	actual := TimeToString(timeTest)
	if actual == "" {
		t.Error("TestTimeToStr result error:", actual)
	}
}

func TestTimestampToStr(t *testing.T) {
	timeStampTest := time.Now().Unix()
	actual := TimestampToString(timeStampTest)
	if actual == "" {
		t.Error("TestTimestampToStr result error:", actual)
	}
}

func TestGetMaxTime(t *testing.T) {
	actual := GetMaxTime()
	if actual.Minute() != time.Now().AddDate(100,0,0).Minute(){
		t.Error("TestTimestampToStr result error:", actual)
	}
}
