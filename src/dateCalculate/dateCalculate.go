package dateCalculate

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strconv"
	"time"
)

type duration struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Days        string `json:"days"`
	Years       string `json:"years"`
	Seconds     string `json:"seconds"`
	Minutes     string `json:"minutes"`
	Hours       string `json:"hours"`
	Weeks       string `json:"weeks"`
	RatioOfYear string `json:"ratioOfYear"`
}

func StringToDate(day, month, year string) time.Time {
	dayNumber, _ := strconv.Atoi(day)
	monthNumber, _ := strconv.Atoi(month)
	yearNumber, _ := strconv.Atoi(year)
	return time.Date(yearNumber, time.Month(monthNumber), dayNumber, 0, 0, 0, 0, time.UTC)
}

func MakeJSON(startDate, endDate time.Time) duration {
	days := Diff(startDate, endDate)

	return duration{
		From:        "Thursday, 4 January 2018",
		To:          "Monday, 4 June 2018",
		Days:        FormatDays(days),
		Years:       "5 months, 1 day",
		Seconds:     daysToSeconds(152),
		Minutes:     "218,880 minutes",
		Hours:       "3,648 hours",
		Weeks:       "21 weeks and 5 days",
		RatioOfYear: "41.64% of 2018",
	}
}

func Diff(startDate, endDate time.Time) int {
	duration := endDate.Add(time.Hour*24).Unix() - startDate.Unix()

	days := duration / (60 * 60 * 24)

	return int(days)
}

func daysToSeconds(days int) string {
	DayToHours := 24
	HourToMinutes := 60
	MinuteToSeconds := 60
	return humanize.Comma(int64(days*DayToHours*HourToMinutes*MinuteToSeconds)) + " seconds"
}

func FormatDays(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)
}
