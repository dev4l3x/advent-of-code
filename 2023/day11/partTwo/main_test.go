package main

import (
	"testing"
)

func TestSumShortestPastBetweenGalaxiesIsCorrect(t *testing.T) {

	space := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	const expectedSum = 8410

	sum := GetSumMinStepsBetweenGalaxies(space)

	if expectedSum != sum {
		t.Fatalf("Expected %v but got %v", expectedSum, sum)	
	}
}
