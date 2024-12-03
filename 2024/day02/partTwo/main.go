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

	reportsWithLevels := make([][]int, 0)

	for _, report := range reports {
		reportParsed := make([]int, 0)
		for _, level := range strings.Split(report, " ") {
			levelNumber := parseNumber(level)
			reportParsed = append(reportParsed, levelNumber)
		}
		reportsWithLevels = append(reportsWithLevels, reportParsed)
	}
	
	safeLevels := 0

	for _, levels := range reportsWithLevels {

		if isSafeReport(levels) {
			safeLevels++
			continue
		}

		fmt.Printf("The report %v is not safe\n", levels)

		for i := 0; i < len(levels); i++ {

			var subreport []int = make([]int, 0)

			if i == 0 {
				subreport = levels[1:]
			} else if i == len(levels) - 1 {
				subreport = levels[:len(levels) - 1]
			} else {
				fmt.Printf("\tLevels: %v\n", levels)
				fmt.Printf("\tGenerating subreport with levels %v, %v\n", levels[:i], levels[i+1:])
				subreport = append(subreport, levels[:i]...)
				subreport = append(subreport, levels[i+1:]...)
			}


			fmt.Printf("\tThe subreport %v is: %v\n", i, subreport)

			if isSafeReport(subreport) {
				safeLevels++
				break
			}
		}	
	}


	fmt.Printf("The number of secure levels is: %d\n", safeLevels)

}

func isSafeReport(levels []int) bool {
	firstLevelNumber := levels[0]
	secondLevelNumber := levels[1]

	if !isSafeDistanceBetweenLevels(firstLevelNumber, secondLevelNumber) {
		return false
	}

	direction := getReportDirection(firstLevelNumber, secondLevelNumber)

	levelIndex := 1

	isLevelSafe := true

	for isLevelSafe == true && levelIndex < len(levels) - 1 {
		levelNumber := levels[levelIndex]
		nextLevelNumber := levels[levelIndex + 1]

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

	return isLevelSafe
}

func isSafeLevel(leftLevel int, rightLevel int, direction int) bool {
	if !isSafeDistanceBetweenLevels(leftLevel, rightLevel) {
		return false
	}

	if direction == UP && leftLevel >= rightLevel {
		return false
	} else if direction == DOWN && leftLevel <= rightLevel {
		return false
	}

	return true
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