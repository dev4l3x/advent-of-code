package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	UP = 1 
	DOWN = -1 
)

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	reports := strings.Split(input, "\n")
	
	safeLevels := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		firstLevelNumber := parseNumber(levels[0])
		secondLevelNumber := parseNumber(levels[1])

		if !isSafeDistanceBetweenLevels(firstLevelNumber, secondLevelNumber) {
			continue
		}

		direction := getReportDirection(firstLevelNumber, secondLevelNumber)

		isLevelSafe := true
		levelIndex := 1

		for isLevelSafe == true && levelIndex < len(levels) - 1 {
			levelNumber := parseNumber(levels[levelIndex])
			nextLevelNumber := parseNumber(levels[levelIndex + 1])

			if !isSafeDistanceBetweenLevels(levelNumber, nextLevelNumber) {
				isLevelSafe = false
				break
			}

			if direction == UP && levelNumber >= nextLevelNumber {
				isLevelSafe = false
			} else if direction == DOWN && levelNumber <= nextLevelNumber {
				isLevelSafe = false
			}

			levelIndex++
		}

		if isLevelSafe {
			safeLevels++
		}
	}


	fmt.Printf("The number of secure levels is: %d\n", safeLevels)

}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}

func getReportDirection(firstLevelNumber int, secondLevelNumber int) int {
	if firstLevelNumber < secondLevelNumber {
		return UP
	}

	return DOWN
}

func isSafeDistanceBetweenLevels(firstLevelNumber int, secondLevelNumber int) bool {
	distance := math.Abs(float64(firstLevelNumber - secondLevelNumber))
	return distance <= 3 && distance >= 1 
}