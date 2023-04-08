package timeutils

import (
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
	DefaultDateFormat = "2006-01-02"
	MinTimStr         = "0001-01-01 00:00:00"
	DateFormat        = "2006-01-02"
)

var ZERO = func() time.Time {
	zero, _ := time.ParseInLocation(DefaultTimeFormat, MinTimStr, time.Local)
	return zero
}()

func Format(t *time.Time, f string) string {
	if t == nil || t.Before(ZERO) || t.Equal(ZERO) {
		return ""
	} else {
		return t.Format(f)
	}
}

func Parse(s string, f string) (time.Time, error) {
	t, err := time.ParseInLocation(f, s, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	if t.Before(ZERO) || t.Equal(ZERO) {
		return ZERO, nil
	} else {
		return time.ParseInLocation(f, s, time.Local)
	}
}

func DefFormat(t *time.Time) string {
	return Format(t, DefaultTimeFormat)
}

func FormatDate(t *time.Time) string {
	return Format(t, DateFormat)
}

func DefParse(s string) (time.Time, error) {
	return Parse(s, DefaultTimeFormat)
}

func ParseDate(s string) (time.Time, error) {
	return Parse(s, DateFormat)
}

func Now() *time.Time {
	tmpTime := time.Now()
	return &tmpTime
}

func GetCurrentDate() time.Time {
	return GetDate(time.Now())
}

func GetDate(t time.Time) time.Time {
	dataDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return dataDate
}

func GetRemainderSecOfCurrentDate() time.Duration {
	now := time.Now()
	ttlh := 23 - now.Hour()
	ttlm := 59 - now.Minute()
	ttls := 59 - now.Second()
	if ttls == 0 {
		ttls = 1
	}
	return time.Duration(ttlh*60*60+ttlm*60+ttls) * time.Second
}
