package dateCalculate

import (
	"testing"
	"time"
)

func Test_MakeJSON_Input_4_1_2018_4_6_2018_Should_be_Struct(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	exception := duration{
		From:        "Thursday, 4 January 2018",
		To:          "Monday, 4 June 2018",
		Days:        "152 days",
		Years:       "5 months, 1 day",
		Seconds:     "13,132,800 seconds",
		Minutes:     "218,880 minutes",
		Hours:       "3,648 hours",
		Weeks:       "21 weeks and 5 days",
		RatioOfYear: "41.64% of 2018",
	}

	actualDuration := MakeJSON(startDate, endDate)

	if actualDuration != exception {
		t.Error("Duration was incorrect, got: ", actualDuration, "want: ", exception)
	}
}

func Test_StringToDate_Input_Day_4_Month_1_Year_2018_Should_be_Date_4_1_2018(t *testing.T) {
	exception := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)

	actualDate := StringToDate("4", "1", "2018")

	if actualDate != exception {
		t.Error("Date was incorrect, got: ", actualDate, "want: ", exception)
	}
}

func Test_Diff_Input_4_1_2018_4_6_2018_Should_be_152(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	exception := 152

	actualDays := Diff(startDate, endDate)

	if actualDays != exception {
		t.Error("Date was incorrect, got: ", actualDays, "want: ", exception)
	}
}

func Test_FormatDate_Input_4_1_2018_Should_be_Thursday_4_January_2018(t *testing.T) {
	date := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	exception := "Thursday, 4 January 2018"

	actualFormatDate := FormatDate(date)
	if actualFormatDate != exception {
		t.Error("Format Date was incorrect, got: ", actualFormatDate, "want: ", exception)
	}
}
func Test_DayToSeconds_Input_152_Should_be_13132800(t *testing.T) {
	days := 152
	exception := "13,132,800 seconds"

	actualSeconds := daysToSeconds(days)

	if actualSeconds != exception {
		t.Error("Date was incorrect, got: ", actualSeconds, "want: ", exception)
	}
}

func Test_FormatDays_Input_152_Should_be_152_days(t *testing.T) {
	days := 152
	exception := "152 days"

	actualDays := FormatDays(days)
	if actualDays != exception {
		t.Error("Days was incorrect, got: ", actualDays, "want: ", exception)
	}
}
