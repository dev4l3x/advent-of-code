package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, error := os.ReadFile("input.txt")	
	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	calibrations := strings.Split(input, "\n")

	var sum int

	for _, calibration := range calibrations {
		calibrationValue := getCalibrationValue(calibration)
		sum += calibrationValue
	}	
	fmt.Printf("The sum of the calibration values is: %v\n", sum)
}

func getCalibrationValue(calibration string) (calibrationValue int) {

	calibrationValues := strings.Split(calibration, "")

	/* numberPattern := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")

	numbers := numberPattern.FindAllString(calibration, -1)
	firstNumber, lastNumber = parseNumber(numbers[0]), parseNumber(numbers[len(numbers) - 1])
	*/

	var numbers []int

	for i := 0; i < len(calibrationValues); i++ {
		number, ok := parseNumber(calibrationValues[i])
		if ok {
			numbers = append(numbers, number)
			continue
		}

		for j := i + 2; j < i + 6 && j <= len(calibration); j++ {
			number, ok := parseNumber(calibration[i:j])
			if ok {
				numbers = append(numbers, number)
				break
			}
		}
	}


	fmt.Println(numbers)
	fmt.Println("Calibration:", calibration, "-", concatNumbers(numbers[0], numbers[len(numbers) - 1]), "-", numbers)

	return concatNumbers(numbers[0], numbers[len(numbers) - 1])
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
