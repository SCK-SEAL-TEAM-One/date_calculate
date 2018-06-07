package dateCalculate

import (
	"testing"
	"time"
)

func Test_MakeJSON_Input_4_1_2018_4_6_2018_Should_Be_Struct(t *testing.T) {
	startDate := "2018-01-04"
	endDate := "2018-06-04"
	exception := duration{
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

	duration := MakeJSON(startDate, endDate)

	if duration != exception {
		t.Error("Duration was incorrect, got: ", duration, "want: ", exception)
	}
}

func Test_StringToDate_Input_Day_4_Month_1_Year_2018_Should_be_Date_4_1_2018(t *testing.T) {
	exception := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)

	actualDate := StringToDate("4", "1", "2018")

	if actualDate != exception {
		t.Error("Date was incorrect, got: ", actualDate, "want: ", exception)
	}
}
