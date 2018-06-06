package dateCalculate

import (
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
)

func StringToTime(year, month, day string) time.Time {
	dayNumber, _ := strconv.Atoi(day)
	monthNumber, _ := strconv.Atoi(month)
	yearNumber, _ := strconv.Atoi(year)

	return time.Date(yearNumber, time.Month(monthNumber), dayNumber, 0, 0, 0, 0, time.UTC)
}
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

func DaysToHours(days int) string {
	hours := days * 24
	numberWithComma := humanize.Comma(int64(hours))

	return numberWithComma + " hours"
}

func FormatDate(date time.Time) string {
	weekDay := date.Weekday().String()
	month := date.Month().String()
	day := date.Day()
	year := date.Year()
	return weekDay + ", " + strconv.Itoa(day) + " " + month + " " + strconv.Itoa(year)
}
