package dateutil

import (
	"fmt"
	"time"
)

func Today() time.Time {
	return time.Now().In(TimeZoneJST)
}

func DateFromYearMonthDay(yyyy int, MM int, dd int) time.Time {
	return time.Date(yyyy, Month(MM), dd, 0, 0, 0, 0, TimeZoneJST)
}

func DateFromDateWithHourMinutes(date time.Time, hour int, minutes int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minutes, 0, 0, TimeZoneJST)
}

func Format24Hour(t time.Time) string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}
