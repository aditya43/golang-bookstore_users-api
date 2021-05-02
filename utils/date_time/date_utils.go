package date_time

import "time"

const dateTimeFormat = "2006-01-02T15:04:05Z"

func GetCurrentDateTime() time.Time {
	return time.Now().UTC()
}

func GetUTCDateTime() string {
	return GetCurrentDateTime().Format(dateTimeFormat)
}
