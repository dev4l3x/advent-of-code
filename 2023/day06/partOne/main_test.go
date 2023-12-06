package main

import (
	"testing"
)

func TestAllPossibleSolutions(t *testing.T) {
	games := map[int]int {
		7: 9,
		15: 40,
		30: 200,
	}

	const expectedProduct = 288

	product := GetProductOfWaysToWinEachGame(games)


	if expectedProduct != product {
		t.Fatalf("Expected %v but got %v", expectedProduct, product)	
	}
}
