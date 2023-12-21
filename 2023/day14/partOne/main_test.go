package main

import (
	"testing"
)

func TestSumRocksLoadIsCorrect(t *testing.T) {

	platform := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	const expectedSum = 136

	sum := GetTotalLoad(platform)

	if expectedSum != sum {
		t.Fatalf("Expected %v but got %v", expectedSum, sum)	
	}
}
