package mbase

// 常用的时间工具.

import (
	"fmt"
	"time"
)

// CurMSecond 当前毫秒数.
func CurMSecond() int64 {
	tm := time.Now()
	return tm.UnixNano() / 1e6
}

// 当前时间 是否 +8 与之前保持一致待定.
func CurTime() time.Time {
	return time.Now()
}

// CurSecond 当前秒数.
func CurSecond() int64 {
	tm := time.Now()
	return tm.Unix()
}

// TimeToDate 返回 yyyy-mm-dd
func TimeToDate() string {
	tm := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", tm.Year(), tm.Month(), tm.Day())
}

// TimeToDBSufix 返回yyyymmdd
func TimeToDBSufix() string {
	tm := time.Now()
	return fmt.Sprintf("%d%02d%02d", tm.Year(), tm.Month(), tm.Day())
}

// TimeToString 返回yyyy-mm-dd hh:mm:ss
func TimeToString() string {
	tm := time.Now()
	return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d",
		tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

//TimeUnixToStr
func TimeUnixToStr(unixSec int64) string {
	tm := time.Unix(unixSec, 0)
	return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d",
		tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

// TimeUnixToDate
func TimeUnixToDate(unixSec int64) time.Time {
	return time.Unix(unixSec, 0)
}

// TimeUnixToDate UTC
func TimeUnixToDateUTC(unixSec int64) time.Time {
	return time.Unix(unixSec, 0).UTC()
}

// 当日第几分钟.
func CurMinOfDay() int {
	tm := time.Now()
	return tm.Hour()*60 + tm.Minute()
}

// StrToTime "yyyy-mm-dd hh:mm:ss"
func StrToTime(strTime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
}

// 超时机制.
func TimeOut(timeout chan bool, times int) {
	go func() {
		time.Sleep(time.Duration(times) * time.Second)
		timeout <- true
	}()
}

func StartTicker(f func(), d time.Duration) chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		timer := time.NewTicker(d)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				f()
			case <-done:
				return
			}
		}
	}()
	return done
}

// 定时器是否已经关闭.
func IsClosed(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
