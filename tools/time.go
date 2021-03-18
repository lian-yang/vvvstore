package tools

import "time"

var (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

// 本月开始时间
func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

// 今日开始时间
func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// 今日结束时间
func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1E9-1, time.Local)
}

// 当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// 当前日期
func NowDate() string {
	return time.Now().Format(TimeLayoutDate)
}

// 当前日期时间
func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

// 解析日期
func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

// 解析日期时间
func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

// 根据字符串解析时间
func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}

