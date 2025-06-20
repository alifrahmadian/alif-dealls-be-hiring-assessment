package handlers

import "time"

func isWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}