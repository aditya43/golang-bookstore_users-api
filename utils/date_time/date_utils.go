package date_time

import "time"

const (
	apiDateTimeFormat = "2006-01-02T15:04:05Z"
	dbDateTimeFormat  = "2006-01-02 15:04:05"
)

func GetCurrentDateTime() time.Time {
	return time.Now().UTC()
}

func GetUTCDateTimeAPIFormat() string {
	return GetCurrentDateTime().Format(apiDateTimeFormat)
}

func GetUTCDateTimeAPIFormatDBFormat() string {
	return GetCurrentDateTime().Format(dbDateTimeFormat)
}
