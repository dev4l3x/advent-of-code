package main

import (
	"testing"
)

func TestCountEnergized(t *testing.T) {

	contraption := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	const expectedSum = 46

	sum := CountEnergized(contraption)

	if expectedSum != sum {
		t.Fatalf("Expected %v but got %v", expectedSum, sum)	
	}
}