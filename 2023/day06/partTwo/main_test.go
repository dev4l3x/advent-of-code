package main

import (
	"testing"
)

func TestAllPossibleSolutions(t *testing.T) {
	badKernedTimes := []int {7, 15, 30}
	badKernedRecords := []int {9, 40, 200}

	const expectedPossibilities = 71503

	possibilities := GetNumberOfWaysToWinGame(badKernedTimes, badKernedRecords)


	if expectedPossibilities != possibilities {
		t.Fatalf("Expected %v but got %v", expectedPossibilities, possibilities)	
	}
}
