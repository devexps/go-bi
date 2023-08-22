package time

import "time"

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

// UnixMilliToString millisecond timestamp -> string
func UnixMilliToString(tm int64) string {
	return time.UnixMilli(tm).Format(TimeLayout)
}
