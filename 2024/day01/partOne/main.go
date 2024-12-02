package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	lines := strings.Split(input, "\n")

	leftNumbers := make([]int, 0)
	rightNumbers := make([]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, "   ")	
		leftNumber := parseNumber(numbers[0]) 
		rightNumber := parseNumber(numbers[1])

		if len(leftNumbers) == 0 {
			leftNumbers = append(leftNumbers, leftNumber)
			rightNumbers = append(rightNumbers, rightNumber)
			continue
		}

		leftIndex := sort.Search(len(leftNumbers), func(i int) bool { return leftNumbers[i] >= leftNumber })
		rightIndex := sort.Search(len(rightNumbers), func(i int) bool { return rightNumbers[i] >= rightNumber })

		leftNumbers = insertNumberAt(leftNumbers, leftNumber, leftIndex)
		rightNumbers = insertNumberAt(rightNumbers, rightNumber, rightIndex)
	}

	totalDistance := 0.0

	for i := 0; i < len(leftNumbers); i++ {
		totalDistance += math.Abs(float64(rightNumbers[i] - leftNumbers[i]))
	}

	fmt.Printf("The total distance is: %f\n", totalDistance)

}

func insertNumberAt(numbers []int, number int, index int) []int {
	if index == len(numbers) {
		return append(numbers, number)
	}

	numbers = append(numbers[:index + 1], numbers[index:]...)
	numbers[index] = number 
	return numbers	
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}
