package main

import "testing"

func TestGameIsPossible(t *testing.T) {
	const game = "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	isPossible := IsGamePossible(game)

	if !isPossible {
		t.Fatalf("Expected game (%v) to be possible but got false instead", game)
	}
}

func TestGameIsNotPossible(t *testing.T) {
	const game = "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"

	isPossible := IsGamePossible(game)

	if isPossible {
		t.Fatalf("Expected game (%v) to be not possible but got true instead", game)
	}
}
