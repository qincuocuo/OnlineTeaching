package mtime

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	zone = time.FixedZone("CST", 8*60*60)
)

const (
	TIME_FORMAT_SECOND           = "2006-01-02 15:04:05"
	TIME_FORMAT_DAY              = "2006-01-02"
	TIME_FORMAT_CHANGE           = "2006-01-02T15:04:05+8:00"
	TIME_FORMAT_NO_DATE          = "15:04:05"
	TIME_FORMAT_SECOND_NO_HYPHEN = "20060102150405"
)

func Zone() *time.Location {
	return zone
}

func SetZone(location *time.Location) {
	zone = location
}

func TimeChangeToString(t time.Time) string {
	return t.In(zone).Format(TIME_FORMAT_CHANGE)
}

func TimeToString(t time.Time) string {
	return t.In(zone).Format(TIME_FORMAT_SECOND)
}

func TimeToNoHyphenString(t time.Time) string {
	return t.In(zone).Format(TIME_FORMAT_SECOND_NO_HYPHEN)
}

func TimeToNoDateString(t time.Time) string {
	return t.In(zone).Format(TIME_FORMAT_NO_DATE)
}

func TimeToDayString(t time.Time) string {
	return t.In(zone).Format(TIME_FORMAT_DAY)
}

func String2Time(rawStr string) (time.Time, error) {
	t, err := time.ParseInLocation(TIME_FORMAT_SECOND, rawStr, zone)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func StringNoHyphen2Time(ctx context.Context, rawStr string) (time.Time, error) {
	t, err := time.ParseInLocation(TIME_FORMAT_SECOND_NO_HYPHEN, rawStr, zone)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func Int64ToTimeString(times int64) string {
	return TimeToString(time.Unix(times, 0))
}

func Int64ToDay(times int64) int64 {
	return int64(time.Unix(times, 0).Sub(time.Now()).Hours() / 24)
}

// Package now is a time toolkit for golang.
//
//  ztime.BeginningOfMinute() // 2013-11-18 17:51:00 Mon
//  ztime.BeginningOfDay()    // 2013-11-18 00:00:00 Mon
//  ztime.EndOfDay()          // 2013-11-18 23:59:59.999999999 Mon

var FirstDayMonday bool
var TimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006-1-2 15:4:5", "2006-1-2 15:4", "2006-1-2", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST"}

type Now struct {
	time.Time
}

func New(t time.Time) *Now {
	return &Now{t}
}

func BeginningOfMinute(t time.Time) time.Time {
	return New(t).BeginningOfMinute()
}

func BeginningOfHour(t time.Time) time.Time {
	return New(t).BeginningOfHour()
}

func BeginningOfDay(t time.Time) time.Time {
	return New(t).BeginningOfDay()
}

func BeginningOfWeek(t time.Time) time.Time {
	return New(t).BeginningOfWeek()
}

func BeginningOfMonth(t time.Time) time.Time {
	return New(t).BeginningOfMonth()
}

func BeginningOfQuarter(t time.Time) time.Time {
	return New(t).BeginningOfQuarter()
}

func BeginningOfYear(t time.Time) time.Time {
	return New(t).BeginningOfYear()
}

func EndOfMinute(t time.Time) time.Time {
	return New(t).EndOfMinute()
}

func EndOfHour(t time.Time) time.Time {
	return New(t).EndOfHour()
}

func EndOfDay(t time.Time) time.Time {
	return New(t).EndOfDay()
}

func EndOfWeek(t time.Time) time.Time {
	return New(t).EndOfWeek()
}

func EndOfMonth(t time.Time) time.Time {
	return New(t).EndOfMonth()
}

func EndOfQuarter(t time.Time) time.Time {
	return New(t).EndOfQuarter()
}

func EndOfYear(t time.Time) time.Time {
	return New(t).EndOfYear()
}

func Monday(t time.Time) time.Time {
	return New(t).Monday()
}

func Sunday(t time.Time) time.Time {
	return New(t).Sunday()
}

func EndOfSunday(t time.Time) time.Time {
	return New(t).EndOfSunday()
}

func Parse(t time.Time, strs ...string) (time.Time, error) {
	return New(t).Parse(strs...)
}

func MustParse(t time.Time, strs ...string) time.Time {
	return New(t).MustParse(strs...)
}

func Between(t time.Time, time1, time2 string) bool {
	return New(t).Between(time1, time2)
}

func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
	return now.Truncate(time.Hour)
}

func (now *Now) BeginningOfDay() time.Time {
	d := time.Duration(-now.Hour()) * time.Hour
	return now.BeginningOfHour().Add(d)
}

func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if FirstDayMonday {
		if weekday == 0 {
			weekday = 7
		}
		weekday = weekday - 1
	}

	d := time.Duration(-weekday) * 24 * time.Hour
	return t.Add(d)
}

func (now *Now) BeginningOfMonth() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.Day())+1) * 24 * time.Hour
	return t.Add(d)
}

func (now *Now) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (now *Now) BeginningOfYear() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.YearDay())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
	return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) Sunday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	} else {
		d := time.Duration(7-weekday) * 24 * time.Hour
		return t.Truncate(time.Hour).Add(d)
	}
}

func (now *Now) EndOfSunday() time.Time {
	return now.Sunday().Add(24*time.Hour - time.Nanosecond)
}

func parseWithFormat(str string) (t time.Time, err error) {
	for _, format := range TimeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

func (now *Now) Parse(strs ...string) (t time.Time, err error) {
	var setCurrentTime bool
	parseTime := []int{}
	currentTime := []int{now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year()}
	currentLocation := now.Location()

	for _, str := range strs {
		onlyTime := regexp.MustCompile(`^\s*\d+(:\d+)*\s*$`).MatchString(str) // match 15:04:05, 15

		t, err = parseWithFormat(str)
		location := t.Location()
		if location.String() == "UTC" {
			location = currentLocation
		}

		if err == nil {
			parseTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
			onlyTime = onlyTime && (parseTime[3] == 1) && (parseTime[4] == 1)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if it is a time only string
				if onlyTime && i <= 2 {
					continue
				}

				// Fill up missed information with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// Default day and month is 1, fill up it if missing it
				if onlyTime {
					if i == 3 || i == 4 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}
		}

		if len(parseTime) > 0 {
			t = time.Date(parseTime[5], time.Month(parseTime[4]), parseTime[3], parseTime[2], parseTime[1], parseTime[0], 0, location)
			currentTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}
	return
}

func (now *Now) MustParse(strs ...string) (t time.Time) {
	t, err := now.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return t
}

func (now *Now) Between(time1, time2 string) bool {
	restime := now.MustParse(time1)
	restime2 := now.MustParse(time2)
	return now.After(restime) && now.Before(restime2)
}

func TimeSub(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Seconds()
}

func CompareTime(careTm time.Time, startTmStr, endTmStr string) int {
	start := fmt.Sprintf("%s %s:00", careTm.Format("2006-01-02"), startTmStr)
	end := fmt.Sprintf("%s %s:00", careTm.Format("2006-01-02"), endTmStr)
	if endTmStr == "24:00" || startTmStr > endTmStr {
		end = fmt.Sprintf("%s 00:00:00", careTm.Add(24*time.Hour))
	}
	startTm, _ := String2Time(start)
	endTm, _ := String2Time(end)
	if careTm.Unix() > startTm.Unix() && careTm.Unix() < endTm.Unix() {
		return 1
	}
	return 0
}

func CheckTime(careTm time.Time, weeks, month []int) string {
	curTm := time.Now()
	curDay := curTm.Day()
	curWeekDay := int(curTm.Weekday())

	monthFunc := func() string {
		for _, day := range month {
			if careTm.Day() == day {
				return ""
			}
		}

		if curDay > month[len(month)-1] {
			return fmt.Sprintf("%d:%d:%d %d:%d:%d",
				curTm.Year(), curTm.Month()+1, month[0], careTm.Hour(), careTm.Minute(), careTm.Second())
		}

		for _, item := range month {
			if item > curDay {
				return fmt.Sprintf("%d:%d:%d %d:%d:%d",
					curTm.Year(), curTm.Month(), item, careTm.Hour(), careTm.Minute(), careTm.Second())
			}
		}
		return ""
	}

	weeksFunc := func() string {
		for _, day := range weeks {
			if careTm.Day() == day {
				return ""
			}
		}

		for _, item := range weeks {
			if item > curWeekDay {
				newDate := curTm.Add(time.Duration(item-curWeekDay) * 24 * time.Hour)
				return fmt.Sprintf("%s %02d:%02d:%02d", newDate.Format("2006-01-02"), careTm.Hour(), careTm.Minute(), careTm.Second())
			}
		}
		return ""
	}

	if len(month) > 0 {
		return monthFunc()
	}

	if len(weeks) > 0 {
		return weeksFunc()
	}

	return ""
}

// 检查是否符合类似这样的格式   00:23
func CheckTimeValidate(tmstrs ...string) bool {
	for _, tmstr := range tmstrs {
		var hour, minute int32
		_, err := fmt.Sscanf(tmstr, "%d:%d", &hour, &minute)
		if err != nil || hour > 24 || hour < 0 || minute >= 60 || minute < 0 {
			return false
		}
	}
	return true
}

func GetNowDiffStr(timeUnit string) (string, error) {
	var err error
	var unitTime time.Duration
	now := time.Now()
	if unitTime, err = time.ParseDuration(timeUnit); err != nil {
		return "", err
	}
	limitTime := now.Add(unitTime).Format("2006-01-02 15:04:05")
	return limitTime, nil
}

func TranStr2Date(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return theTime
}
