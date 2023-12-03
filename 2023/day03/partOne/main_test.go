package main

import (
	"testing"
)

func TestPartNumberSumIsCorrect(t *testing.T) {
	engine := "467..114..\n...%......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

	partNumbersSum := GetPartNumbersSumFromEngine(engine)

	if partNumbersSum != 4361 {
		t.Fatalf("Expected part numbers sum of 4361 but got %v", partNumbersSum)
	} 
}

func TestPartNumberSumIsCorrect2(t *testing.T) {
	engine := "467..114..\n...%......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592&1...\n......755.\n...$.*....\n.664.598.."

	partNumbersSum := GetPartNumbersSumFromEngine(engine)

	if partNumbersSum != 4362 {
		t.Fatalf("Expected part numbers sum of 4361 but got %v", partNumbersSum)
	} 
}
