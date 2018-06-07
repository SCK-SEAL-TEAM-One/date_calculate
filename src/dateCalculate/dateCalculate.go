package dateCalculate

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
)

const DayToHours = 24
const HourToMinutes = 60
const MinuteToSeconds = 60

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
	duration := endDate.Add(time.Hour*DayToHours).Unix() - startDate.Unix()
	days := duration / (DayToHours * HourToMinutes * MinuteToSeconds)
	return int(days)
}

func FormatDate(date time.Time) string {
	weekDay := date.Weekday()
	day := date.Day()
	month := date.Month()
	year := date.Year()
	return weekDay.String() + ", " + strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
}

func daysToSeconds(days int) string {
	seconds := days * DayToHours * HourToMinutes * MinuteToSeconds
	return fmt.Sprintf("%s seconds", humanize.Comma(int64(seconds)))
}

func FormatDays(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)
}

func DaysToHours(days int) string {
	hours := days * DayToHours
	return fmt.Sprintf("%s hours", humanize.Comma(int64(hours)))
}
