package dateCalculate

import (
	"testing"
	"time"
)

func Test_Diff_Input_4_1_2018_4_6_2018_should_be_152(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)

	expectedResult := 152
	duration := Diff(startDate, endDate)
	if expectedResult != duration {
		t.Error("Expected: ", expectedResult, " but got ", duration)
	}
}
