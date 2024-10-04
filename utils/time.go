package utils

import (
	"time"
)

type TimeCountdown struct {
	Total   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func GetRemainingTime(currentTime, timeout time.Time) TimeCountdown {
	difference := timeout.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return TimeCountdown{
		Total:   total,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}

func RangeDate(startDate, endDate time.Time) func() time.Time {
	y, m, d := startDate.Date()
	startDate = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = endDate.Date()
	endDate = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if startDate.After(endDate) {
			return time.Time{}
		}
		date := startDate
		startDate = startDate.AddDate(0, 0, 1)
		return date
	}
}

func DateIsEqualOrAfter(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 >= y2 && m1 >= m2 && d1 >= d2
}

func InTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) && check.Before(end)) || check == start || check == end
}

func NullTimeScan(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
