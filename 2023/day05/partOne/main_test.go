package main

import (
	"testing"
	"os"
	"strings"
)

func TestLowerLocationIsRetrieved(t *testing.T) {
	file, err := os.ReadFile("testData/test1.txt")	
	if err != nil {
		t.Fatalf("An error has ocurred while reading input")	
	}

	input := strings.Split(string(file), "\n\n")

	const expectedLocation = 35
	lowestLocation := GetLowestLocation(input)

	if expectedLocation != lowestLocation {
		t.Fatalf("Expected %v but got %v", expectedLocation, lowestLocation)	
	}
}
