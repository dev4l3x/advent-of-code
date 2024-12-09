package main

import (
	"fmt"
	"log"
	"math"
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

	nextElement := calibration.equation[nextElementIndex]
	sum := previousComputedValue + nextElement
	mul := previousComputedValue * nextElement
	concat := concat(previousComputedValue, nextElement)

	sumBranchSatisfies := false
	mulBranchSatisifes := false
	concatBranchSatisifies := false

	if sum <= calibration.value {
		sumBranchSatisfies = isTrueCalibration(sum, nextElementIndex + 1, calibration)
	}

	if mul <= calibration.value {
		mulBranchSatisifes = isTrueCalibration(mul, nextElementIndex + 1, calibration)
	}

	if concat <= calibration.value {
		concatBranchSatisifies = isTrueCalibration(concat, nextElementIndex + 1, calibration)
	}

	return sumBranchSatisfies || mulBranchSatisifes || concatBranchSatisifies
}

func concat(leftOperator int, rightOperator int) int {
	rightOperatorDigits := int(math.Log10(float64(rightOperator)) + 1)

	shiftedLeftOperator := leftOperator * int(math.Pow(10, float64(rightOperatorDigits)))

	return shiftedLeftOperator + rightOperator	
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}