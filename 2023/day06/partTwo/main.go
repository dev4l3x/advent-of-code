package main

import (
	"fmt"
	"math"
)

func main() {
	badKernedTimes := []int {51, 69, 98, 78}
	badKernedRecords := []int {377, 1171, 1224, 1505}


	product := GetNumberOfWaysToWinGame(badKernedTimes, badKernedRecords)

	fmt.Println("The total product is:", product)
}


func getGame(badKernedTimes []int, badKernedRecords []int) (totalTime int, totalRecord int) {

	totalTime = badKernedTimes[0]
	totalRecord = badKernedRecords[0]

	for i := 1 ; i < len(badKernedTimes) ; i++ {
		totalTime = concatNumbers(totalTime, badKernedTimes[i])
		totalRecord = concatNumbers(totalRecord, badKernedRecords[i])
	}

	return totalTime, totalRecord
}

func concatNumbers(left int, right int) int {
	position := len(fmt.Sprint(right))
	factor := math.Pow10(position)
	return (int(factor) * left) + right
}

func GetNumberOfWaysToWinGame(badKernedTime []int, badKernedRecord []int) int {
	const INCREASING_SPEED = 1

	time, currentRecord := getGame(badKernedTime, badKernedRecord)
	fmt.Println("Time:", time, "; Record:", currentRecord)

	var numberOfWays int

	for holdTime := 1 ; holdTime < time ; holdTime++ {
		totalSpeed := holdTime	
		leftTime := time - holdTime
		totalDistance := leftTime * totalSpeed
		if currentRecord < totalDistance {
			numberOfWays++
		}
	}
	
	return numberOfWays
}
