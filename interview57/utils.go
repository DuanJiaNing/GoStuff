package interview57

import (
	"./model"
	"math"
	"time"
)

// month (January = 1, ...)
func getDate(year int, month time.Month, day int, hour int, minute int, second int) time.Time {
	return time.Date(year, month, day, hour, minute, second, 0, time.Local)
}

func calcTimeDurationInDays(interval model.TimeInterval) int {
	return int(math.Floor(float64(interval.EndTime.Day() - interval.StartTime.Day())))
}
