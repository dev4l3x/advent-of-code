package main

import (
	"strings"
	"testing"
)

func TestNumberOfStepsInNetwork(t *testing.T) {

	report := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	const expectedExtrapolatedValues = 114

	extrapolatedValues := GetSumExtrapolatedValues(strings.Split(report, "\n"))

	if expectedExtrapolatedValues != extrapolatedValues {
		t.Fatalf("Expected %v but got %v", expectedExtrapolatedValues, extrapolatedValues)	
	}
}
