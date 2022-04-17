package mbase

import (
	"fmt"
	"time"
)

// TimeToDate 返回 [yyyy-mm-dd]
func TimeToDate(tm time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", tm.Year(), tm.Month(), tm.Day())
}

// TimeToDBSuffix 返回 [yyyymmdd]
func TimeToDBSuffix(tm time.Time) string {
	return fmt.Sprintf("%d%02d%02d", tm.Year(), tm.Month(), tm.Day())
}

// TimeToStr 时间转化为字符串
func TimeToString(tm time.Time) (timeStr string) {
	if tm.IsZero() {
		return ""
	}
	location, _ := time.LoadLocation("Local")
	timeStr = tm.In(location).Format("2006-01-02 15:04:05")
	return
}

// TimestampToStr 时间戳转化为字符串
func TimestampToString(stamp int64) (timeStr string) {
	timeUnix := time.Unix(stamp, 0)
	location, _ := time.LoadLocation("Local")
	timeStr = timeUnix.In(location).Format("2006-01-02 15:04:05")
	return
}

// GetMaxTime 获取最大时间100年后(默认永不过期)
func GetMaxTime() (timeStr time.Time) {
	return time.Now().AddDate(100, 0, 0)
}
