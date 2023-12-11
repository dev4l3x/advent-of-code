package main

import (
	"testing"
)

func TestNumberOfStepsInNetwork(t *testing.T) {

	pipesMap := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	const expectedFarthestPointSteps = 8

	farthestPointSteps := GetStepsFromFarthestPoint(pipesMap)

	if expectedFarthestPointSteps != farthestPointSteps {
		t.Fatalf("Expected %v but got %v", expectedFarthestPointSteps, farthestPointSteps)	
	}
}
