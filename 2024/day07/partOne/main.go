package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calibration struct{
	value int
	equation []int
}

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	calibrations := strings.Split(input, "\n")

	parsedCalibrations := make([]Calibration, 0)

	for _, calibration := range calibrations {
		parsedCalibration := strings.Split(calibration, ": ")

		value := parseNumber(parsedCalibration[0])
		equation := make([]int, 0)

		stringEquation := strings.Split(parsedCalibration[1], " ")
		for _, term := range stringEquation {
			equation = append(equation, parseNumber(term))
		}

		parsedCalibrations = append(parsedCalibrations, Calibration{value, equation})
	}

	sum := 0

	for _, calibration := range parsedCalibrations {
		if isTrueCalibration(calibration.equation[0], 1, calibration) {
			sum += calibration.value
		}
	}

	fmt.Printf("Sum true calibrations: %v\n", sum)
}

func isTrueCalibration(previousComputedValue int, nextElementIndex int, calibration Calibration) bool {

	if nextElementIndex == len(calibration.equation) {
		return previousComputedValue == calibration.value
	}

	sum := previousComputedValue + calibration.equation[nextElementIndex]
	mul := previousComputedValue * calibration.equation[nextElementIndex]

	sumBranchSatisfies := false
	mulBranchSatisifes := false

	if sum <= calibration.value {
		sumBranchSatisfies = isTrueCalibration(sum, nextElementIndex + 1, calibration)
	}

	if mul <= calibration.value {
		mulBranchSatisifes = isTrueCalibration(mul, nextElementIndex + 1, calibration)
	}

	return sumBranchSatisfies || mulBranchSatisifes 
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}