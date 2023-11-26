package main

import (
	"testing"
)

func TestSumThreeElfsWithMostKcal(t *testing.T) {

	sum := getSumOfElfsWithMostKcal()
	expectedSum := 201491

	if sum !=  expectedSum {
		t.Fatalf("Expected %v to be equal to %v", sum, expectedSum)
	}

}
