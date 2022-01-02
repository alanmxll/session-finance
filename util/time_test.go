package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2022-01-02T10:53:25")

	if convertedTime.Year() != 2022 {
		t.Errorf("The year should be 2022")
	}

	if convertedTime.Month() != 01 {
		t.Errorf("The month should be 01")
	}

	if convertedTime.Day() != 02 {
		t.Errorf("The day should be 02")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("The hour should be 10")
	}

	if convertedTime.Minute() != 53 {
		t.Errorf("The minute should be 53")
	}

	if convertedTime.Second() != 25 {
		t.Errorf("The second should be 25")
	}
}
