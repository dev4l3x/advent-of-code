package main

import (
	"testing"
)

func TestGearRatioSumIsCorrect(t *testing.T) {
	engine := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

	gearRatioSum := GetPartNumbersSumFromEngine(engine)

	if gearRatioSum != 467835{
		t.Fatalf("Expected part numbers sum of 467835 but got %v", gearRatioSum)
	} 
}
