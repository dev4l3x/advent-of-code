package main

import (
	"testing"
)

func TestNumberOfStepsInNetwork(t *testing.T) {

	network := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	const expectedSteps = 6

	steps := GetStepsFromNetwork(network)

	if expectedSteps != steps {
		t.Fatalf("Expected %v but got %v", expectedSteps, steps)	
	}
}
