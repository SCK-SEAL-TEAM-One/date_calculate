package dateCalculate

import (
	"strconv"
	"time"
)

type duration struct {
	Form        string
	To          string
	Days        string
	Years       string
	Seconds     string
	Minutes     string
	Hours       string
	Weeks       string
	RatioOfYear string
}

func StringToDate(day, month, year string) time.Time {
	dayNumber, _ := strconv.Atoi(day)
	monthNumber, _ := strconv.Atoi(month)
	yearNumber, _ := strconv.Atoi(year)
	return time.Date(yearNumber, time.Month(monthNumber), dayNumber, 0, 0, 0, 0, time.UTC)
}

func MakeJSON(startDate, endDate time.Time) duration {
	return duration{
		Form:        "Thursday, 4 January 2018",
		To:          "Monday, 4 June 2018",
		Days:        "152 days",
		Years:       "5 months, 1 days",
		Seconds:     "13,132,800 seconds",
		Minutes:     "218,880 minutes",
		Hours:       "3,648 hours",
		Weeks:       "1 weeks 7 days",
		RatioOfYear: "41.64% of 2018",
	}
}
