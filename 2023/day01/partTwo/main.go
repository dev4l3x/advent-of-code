package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {
	file, error := os.ReadFile(fileName)	
	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	return strings.Split(input, "\n")
}

func main() {
	calibrations := readInput(os.Args[1])	

	calibrationsSum := getCalibrations(calibrations)
		
	fmt.Printf("The sum of the calibration values is: %v\n", calibrationsSum)
}

func getCalibrations(calibrations []string) (sum int) {
	for _, calibration := range calibrations {
		calibrationValue := getCalibrationValue(calibration)
		sum += calibrationValue
	}
	return sum
}

func getCalibrationValue(calibration string) (calibrationValue int) {

	calibrationValues := strings.Split(calibration, "")

	var numbersInCalibration []int

	// A minor improvement would be to just take the first and the last items in the calibration to avoid parsing all the numbers
	for i := 0; i < len(calibrationValues); i++ {
		number, ok := parseNumber(calibrationValues[i])
		if ok {
			numbersInCalibration = append(numbersInCalibration, number)
			continue
		}

		const MAX_CHARS_TEXTUAL_NUMBER = 5
		const MIN_CHARS_TEXTUAL_NUMBER = 3
		for j := i + MIN_CHARS_TEXTUAL_NUMBER - 1; j <= i + MAX_CHARS_TEXTUAL_NUMBER && j < len(calibration); j++ {
			number, ok := parseNumber(calibration[i:j+1])
			if ok {
				numbersInCalibration = append(numbersInCalibration, number)
				break
			}
		}
	}

	return concatNumbers(numbersInCalibration[0], numbersInCalibration[len(numbersInCalibration) - 1])
}

func concatNumbers(first int, last int) (concatedNumbers int) {
	return first * 10 + last
}

var textualNumbers = map[string]int {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func parseNumber(number string) (int, bool) {
	
	textualNumber, ok := textualNumbers[number]

	if (ok) {
		return textualNumber, true
	}

	parsedNumber, error := strconv.Atoi(number)

	if (error != nil) {
		return 0, false
	}

	return parsedNumber, true
}
