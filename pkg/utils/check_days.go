package utils

import "time"

func IsWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}

func IsWeekday(date time.Time) bool {
	day := date.Weekday()
	return day != time.Saturday && day != time.Sunday
}