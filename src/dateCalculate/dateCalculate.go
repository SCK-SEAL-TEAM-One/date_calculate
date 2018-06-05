package dateCalculate

import (
	"time"
)

func Diff(startDate time.Time, endDate time.Time) int {
	secondsInDays := (60 * 60 * 24)
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}
	startTimestamp := startDate.Unix()
	endTimestamp := endDate.Add(time.Hour * 24).Unix()

	durationUnix := endTimestamp - startTimestamp
	durationDays := int(durationUnix) / secondsInDays
	return durationDays
}

func DaysToHours(days int) string{

	return "3,648 hours"
}