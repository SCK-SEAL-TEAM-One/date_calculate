package dateCalculate

import (
	"testing"
	"time"
)

func Test_Diff_Input_4_1_2018_4_6_2018_Should_be_152(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)

	expectedResult := 152
	duration := Diff(startDate, endDate)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
	}
}

func Test_DaysToHours_Input_152_Should_be_3648_hours(t *testing.T) {
	days := 152

	expectedResult := "3,648 hours"
	duration := DaysToHours(days)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
	}
}

func Test_DaysToMinutes_Input_152_Should_be_218880_minutes(t *testing.T) {
	days := 152

	expectedResult := "218,880 minutes"
	duration := DaysToMinutes(days)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
	}
}

func Test_DaysToSeconds_Input_152_Should_be_13132800_seconds(t *testing.T){
	days := 152

	expectedResult := "13,132,800 seconds"
	duration := DaysToSeconds(days)
	if expectedResult != duration {
		t.Error("Expected: ",expectedResult, " but got ", duration)
	}
}

func Test_DaysToWeeks_Input_152_Should_be_21_weeks_and_5_days_seconds(t *testing.T){
	days := 152

	expectedResult := "21 weeks and 5 days"
	duration := DaysToWeeks(days)
	if expectedResult != duration {
		t.Error("Expected: ",expectedResult, " but got ", duration)
	}
}

func Test_DaysToRatioOfYears_Input_152_4_1_2018_4_6_2018_Should_be_41_dot_64_percent_of_2018(t *testing.T){
	days := 152
	inputStartDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	inputEndDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	expectedResult := "41.64% of 2018"
	duration := DaysToRatioOfYears(days,inputStartDate,inputEndDate)
	if expectedResult != duration {
		t.Error("Expected: ",expectedResult, " but got ", duration)
	}
}

func Test_DaysToRatioOfYears_Input_152_4_1_2018_4_6_2025_Should_be_742_dot_19_percent_of_common_year_365_days(t *testing.T){
	days := 2709
	inputStartDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	inputEndDate := time.Date(2025, 6, 4, 0, 0, 0, 0, time.UTC)
	expectedResult := "742.19% of common year (365 days)"
	duration := DaysToRatioOfYears(days,inputStartDate,inputEndDate)
	if expectedResult != duration {
		t.Error("Expected: ",expectedResult, " but got ", duration)
	}
}

func Test_FormatDate_Input_4_1_2018_Should_be_Thursday_4_January_2018(t *testing.T) {
	inputDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	outputDateFormat := FormatDate(inputDate)

	expectedResult := "Thursday, 4 January 2018"
	if expectedResult != outputDateFormat {
		t.Error("Expected: ", expectedResult, " but got ", outputDateFormat)
	}
}

func Test_FormatDate_Input_4_6_2018_Should_be_Monday_4_June_2018(t *testing.T) {
	inputDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	outputDateFormat := FormatDate(inputDate)

	expectedResult := "Monday, 4 June 2018"
	if expectedResult != outputDateFormat {
		t.Error("Expected: ", expectedResult, " but got ", outputDateFormat)
	}
}

func Test_StringToTime_Input_4_6_2018_Should_be_Monday_4_June_2018(t *testing.T) {
	outputDate := StringToTime("2018", "6", "4")
	expectedResult := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	if expectedResult != outputDate {
		t.Error("Expected: ", expectedResult, " but got ", outputDate)
	}
}
