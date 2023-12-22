package main

import (
	"testing"
)

func TestCalculateFocusingPower(t *testing.T) {

	platform := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	const expectedSum = 145

	sum := CalculateFocusingPower(platform)

	if expectedSum != sum {
		t.Fatalf("Expected %v but got %v", expectedSum, sum)	
	}
}
