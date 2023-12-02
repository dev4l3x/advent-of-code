package main

import "testing"

func TestMinimumNumberOfCubes(t *testing.T) {
	const game = "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	r, g, b := GetMinimumNumberOfCubes(game)

	if r != 4 {
		t.Fatalf("Expected r to be 4")
	}
	if g != 2 {
		t.Fatalf("Expected g to be 2")
	}
	if b != 6 {
		t.Fatalf("Expected b to be 6")
	}
}
