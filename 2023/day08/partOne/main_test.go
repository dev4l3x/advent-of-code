package main

import (
	"testing"
)

func TestNumberOfStepsInNetwork(t *testing.T) {

	network := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	const expectedSteps = 6

	steps := GetStepsFromNetwork(network)

	if expectedSteps != steps {
		t.Fatalf("Expected %v but got %v", expectedSteps, steps)	
	}
}
