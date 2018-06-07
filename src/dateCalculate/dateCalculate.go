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
const weekToDays = 7

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
		From:        FormatDate(startDate),
		To:          FormatDate(endDate),
		Days:        FormatDays(days),
		Years:       "5 months, 1 day",
		Seconds:     DaysToSeconds(days),
		Minutes:     DaysToMinutes(days),
		Hours:       DaysToHours(days),
		Weeks:       DaysToWeeks(days),
		RatioOfYear: RatioOfYear(days, startDate, endDate),
	}
}

func Diff(startDate, endDate time.Time) int {
	duration := endDate.Add(time.Hour*DayToHours).Unix() - startDate.Unix()
	days := duration / (DayToHours * HourToMinutes * MinuteToSeconds)
	return int(days)
}

func FormatDate(date time.Time) string {
	weekDay := date.Weekday().String()
	day := strconv.Itoa(date.Day())
	month := date.Month().String()
	year := strconv.Itoa(date.Year())
	return fmt.Sprintf("%s, %s %s %s", weekDay, day, month, year)
}

func DaysToSeconds(days int) string {
	seconds := days * DayToHours * HourToMinutes * MinuteToSeconds
	return fmt.Sprintf("%s seconds", humanize.Comma(int64(seconds)))
}

func DaysToMinutes(days int) string {
	hours := days * DayToHours * HourToMinutes
	return fmt.Sprintf("%s minutes", humanize.Comma(int64(hours)))
}

func DaysToHours(days int) string {
	hours := days * DayToHours
	return fmt.Sprintf("%s hours", humanize.Comma(int64(hours)))
}

func RatioOfYear(days int, startDate, endDate time.Time) string {
	daysInYears := 365
	if startDate.Year() == endDate.Year() {
		ratioOfYear := float64(days) / float64(daysInYear(startDate.Year())) * 100
		return fmt.Sprintf("%.2f%% of %d", ratioOfYear, startDate.Year())
	}
	ratioOfYear := float64(days) / float64(daysInYears)
	return fmt.Sprintf("%.2f%% of a common year (365 days)", ratioOfYear)
}

func FormatDays(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%s days", humanize.Comma(int64(days)))
}

func daysInYear(year int) int {
	daysInFebruary := time.Date(year, time.February+1, 0, 0, 0, 0, 0, time.UTC).Day()
	return daysInFebruary + 365 - 28
}

func DaysToWeeks(days int) string {
	weeks := days / weekToDays
	weeksDays := days % weekToDays

	return fmt.Sprintf("%d weeks and %d days", weeks, weeksDays)
}
