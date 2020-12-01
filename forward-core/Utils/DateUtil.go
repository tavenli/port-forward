package Utils

import (
	"fmt"
	"math"
	"time"
)

//  GO的诞辰
const timeLayout = "2006-01-02 15:04:05"

//  取当前系统时间
func GetTimeNow() time.Time {
	return time.Now()
}

func GetTime(timeStr string) time.Time {
	toTime, _ := ToTime(timeStr)
	return toTime
}

func JavaLongTime(javaLong int64) time.Time {
	//1492566520958	-> 2017-04-19 09:48:40
	//fmt.Println(time.Unix(1492566520958/1000, 0))
	//fmt.Println(time.Unix(0, 1492566520958*1000000))
	return time.Unix(0, javaLong*1000000)
}

func ToTime(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	toTime, err := time.ParseInLocation(timeLayout, timeStr, loc)
	//toTime, err := time.Parse(timeLayout, timeStr)
	return toTime, err

}

func ToTimeByFm(timeStr string, format string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	toTime, err := time.ParseInLocation(format, timeStr, loc)
	//toTime, err := time.Parse(timeLayout, timeStr)
	return toTime, err

}

//要想格式化为：yyyyMMddHHmmss
//则 format = "20060102150405"
//要想格式化为：yyyy-MM-dd HH:mm:ss
//则 format = "2006-01-02 15:04:05"
//要想格式化为：yyyy-MM-dd
//则 format = "2006-01-02"
func FormatTimeByFm(t time.Time, format string) string {

	return t.Format(format)
}

func GetCurrentTime() string {
	return FormatTime(time.Now())
}

func GetCurrentDay() string {
	return FormatTimeByFm(time.Now(), "2006-01-02")
}

func FormatTime(t time.Time) string {
	//
	return FormatTimeByFm(t, "2006-01-02 15:04:05")
}

func FormatTimeToNum(t time.Time) string {
	//
	return FormatTimeByFm(t, "20060102150405")
}

//  在当前时间之前
func IsBeforeNow(t time.Time) (result bool) {
	result = false
	if &t != nil && t.Before(time.Now()) {
		result = true
	}
	return
}

//  在当前时间之后
func IsAfterNow(t time.Time) (result bool) {
	result = false
	if &t != nil && t.After(time.Now()) {
		result = true
	}
	return
}

func SubDateTime(firstTime time.Time, secondTime time.Time) (result time.Duration) {
	result = time.Duration(0)
	if &firstTime != nil && &secondTime != nil {
		result = secondTime.Sub(firstTime)
	}
	return
}

func DifferDays(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Hours()
	return int64(math.Abs(result) / 24)
}

func DifferHour(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Hours()
	//return int64(result) 两个时间的先后顺序不一样，可能出现负数
	return int64(math.Abs(result))
}

func DifferMin(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Minutes()
	return int64(math.Abs(result))
}

func DifferSec(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Seconds()
	return int64(math.Abs(result))
}

//  24小时前的时间
func Before24h() time.Time {
	t, _ := time.ParseDuration("-24h")
	return time.Now().Add(t)
}

func AddSecs(_time time.Time, secs int64) time.Time {
	t, _ := time.ParseDuration("1s")
	return time.Now().Add(t * time.Duration(secs))
}

/*
   增加10分钟：utils.AddMins(time.Now(), 10)
   减少5分钟：utils.AddMins(time.Now(), -5)
*/
func AddMins(_time time.Time, mins int64) time.Time {
	t, _ := time.ParseDuration("1m")
	return time.Now().Add(t * time.Duration(mins))
}

func AddHours(_time time.Time, hours int64) time.Time {
	t, _ := time.ParseDuration("1h")
	return time.Now().Add(t * time.Duration(hours))
}

func AddDays(_time time.Time, days int) time.Time {
	return _time.AddDate(0, 0, days)
}

func AddMonths(_time time.Time, months int) time.Time {
	return _time.AddDate(0, months, 0)
}

func GetBeginTime(_time time.Time) time.Time {
	//2017-06-28 00:00:00 +0800 CST
	return GetBeginTimeByLoc(_time, time.Local)
	//return GetBeginTimeByLoc(_time, time.UTC)

}

func GetEndTime(_time time.Time) time.Time {
	//2017-06-28 23:59:59.999999999 +0800 CST
	return GetEndTimeByLoc(_time, time.Local)
}

func GetBeginTimeByLoc(_time time.Time, loc *time.Location) time.Time {
	year, month, day := _time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)

}

func GetEndTimeByLoc(_time time.Time, loc *time.Location) time.Time {
	year, month, day := _time.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, loc)
}

// 一行代码计算代码执行时间
// defer utils.TimeCost(time.Now())
func TimeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println("TimeCost：", terminal)
}
