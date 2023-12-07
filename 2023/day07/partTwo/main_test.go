package main

import (
	"testing"
)

func TestTotalWinningsIsCorrect(t *testing.T) {

	hands := []string {"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}

	const expectedTotalWinnings = 5905

	totalWinnings := GetTotalWinningsFromHands(hands)

	if expectedTotalWinnings != totalWinnings {
		t.Fatalf("Expected %v but got %v", expectedTotalWinnings, totalWinnings)	
	}
}
