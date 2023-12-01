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

	values := strings.Split(calibration, "")	

	var firstNumber, lastNumber string

	for _, value := range values {
		if isNumber(value) {
			firstNumber = value
			break
		}
	}

	for end := len(values) - 1 ; end >= 0 ; end-- {
		if isNumber(values[end]) {
			lastNumber = values[end]
			break
		}
	}

	number, err := strconv.Atoi(firstNumber + lastNumber)
	if err != nil {
		fmt.Printf("An error has ocurred while parsing value in calibration (%v): %v", calibration, err)
		os.Exit(1)
	}

	return number
}

func isNumber(number string) bool {
	_, error := strconv.Atoi(number)

	return error == nil
}
