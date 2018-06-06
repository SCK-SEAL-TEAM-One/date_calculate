package dateCalculate

import (
	"testing"
	"time"
)

func Test_Diff_Input_4_1_2018_4_6_2018_should_be_152(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)

	expectedResult := 152
	duration := Diff(startDate, endDate)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
	}
}

func Test_DaysToHours_Input_152_should_be_3648_hours(t *testing.T) {
	days := 152

	expectedResult := "3,648 hours"
	duration := DaysToHours(days)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
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
