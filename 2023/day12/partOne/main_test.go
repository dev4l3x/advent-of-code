package main

import (
	"testing"
)

func TestGetSumPossibleArrangementsIsCorrect(t *testing.T) {

	space := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

	const expectedSum = 21

	sum := GetSumPossibleArrangements(space)

	if expectedSum != sum {
		t.Fatalf("Expected %v but got %v", expectedSum, sum)	
	}
}
